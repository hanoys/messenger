package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/hanoy/messenger/internal/config"
)

var (
	tokenExpiredErr = errors.New("token expired")
	invalidTokenErr = errors.New("invalid token")
)

type Payload struct {
	ID     uuid.UUID
	UserID int
}

func NewPayload(userID int) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:     tokenID,
		UserID: userID,
	}, nil
}

type JWTClaims struct {
	Payload
	jwt.StandardClaims
}

type Provider struct {
	cfg *config.Config
}

func NewProvider(config *config.Config) *Provider {
	return &Provider{config}
}

func (p *Provider) CreateToken(payload *Payload) (string, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(p.cfg.JWT.TokenExpirationTime))
	claims := &JWTClaims{
		Payload: *payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(p.cfg.JWT.SecretKey))
	return tokenString, err
}

func (p *Provider) VerifyToken(tokenString string) (*Payload, error) {
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
    log.Println("token user id: ", claims.Payload.UserID)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, tokenExpiredErr
	}
	return &claims.Payload, nil
}
