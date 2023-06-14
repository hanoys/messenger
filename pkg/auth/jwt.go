package auth

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/hanoy/messenger/internal/config"
	"github.com/redis/go-redis/v9"
)

var (
	tokenExpiredErr = errors.New("token expired")
	invalidTokenErr = errors.New("invalid token")
)

type TokenSession struct {
	SessionID      uuid.UUID
	Tokens         *TokenPair
	ExpirationTime time.Time
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Payload struct {
	SessionID uuid.UUID
	UserID    int
	Role      string
}

func NewPayload(userID int, role string) (*Payload, error) {
	sessionID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		SessionID: sessionID,
		UserID:    userID,
		Role:      role,
	}, nil
}

type JWTClaims struct {
	Payload
	jwt.RegisteredClaims
}

type Provider struct {
	redisClient *redis.Client
	cfg         *config.Config
}

func NewProvider(redisClient *redis.Client, cfg *config.Config) *Provider {
	return &Provider{redisClient: redisClient,
		cfg: cfg}
}

func (p *Provider) newTokenWithExpiration(ctx context.Context, payload *Payload, exp time.Time) (string, error) {
	claims := &JWTClaims{
		Payload: *payload,
		RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(exp),
        },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(p.cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (p *Provider) NewSession(ctx context.Context, payload *Payload) (*TokenSession, error) {
	accessExpTime := time.Now().Add(time.Minute * time.Duration(p.cfg.JWT.AccessTokenExpTime))
	refreshExpTime := time.Now().Add(time.Minute * time.Duration(p.cfg.JWT.RefreshTokenExpTime))

	accessTokenString, err := p.newTokenWithExpiration(ctx, payload, accessExpTime)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := p.newTokenWithExpiration(ctx, payload, refreshExpTime)
	if err != nil {
		return nil, err
	}

	session := &TokenSession{
		SessionID: payload.SessionID,
		Tokens: &TokenPair{
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString},
		ExpirationTime: refreshExpTime,
	}

	tokensJSON, err := json.Marshal(session.Tokens)
	if err != nil {
		return nil, err
	}

	_, err = p.redisClient.Set(ctx, session.SessionID.String(),
		string(tokensJSON), session.ExpirationTime.Sub(time.Now())).Result()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (p *Provider) RefreshSession(ctx context.Context, refreshTokenString string) (*TokenSession, error) {
	refreshClaims, err := p.parseToken(refreshTokenString)
	if err != nil {
		return nil, err
	}

	p.redisClient.Del(ctx, refreshClaims.Payload.SessionID.String()).Result()

	return p.NewSession(ctx, &refreshClaims.Payload)
}

func (p *Provider) CloseSession(ctx context.Context, tokenString string) error {
	claims, err := p.parseToken(tokenString)
	if err != nil {
		return err
	}

	ok, err := p.redisClient.Del(ctx, claims.Payload.SessionID.String()).Result()
	if err != nil {
		return err
	}

	if ok != 1 {
		return errors.New("tokent wasn't deleted")
	}

	return nil
}

func (p *Provider) VerifyToken(ctx context.Context, tokenString string) (*Payload, error) {
	claims, err := p.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, tokenExpiredErr
	}

	_, err = p.redisClient.Get(ctx, claims.Payload.SessionID.String()).Result()
	if err != nil {
		return nil, err
	}

	return &claims.Payload, nil
}

func (p *Provider) parseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(p.cfg.JWT.SecretKey), nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, invalidTokenErr
	}

	claims := token.Claims.(*JWTClaims)

	return claims, nil
}

