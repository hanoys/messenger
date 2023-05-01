package app

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type Logger struct {
	handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s\n", r.Host)
	log.Printf("URL: %s\n", r.URL)
	log.Printf("Method: %s\n", r.Method)

	log.Printf("Cookies:\n")
	for _, c := range r.Cookies() {
		log.Println(c)
	}

	l.handler.ServeHTTP(w, r)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	login_path := filepath.Join("static", "login.html")
	templ := template.Must(template.ParseFiles(login_path))
	templ.Execute(w, nil)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&user)
	log.Printf("user: %v\n", user)
	w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(user)
	http.SetCookie(w, &http.Cookie{Name: "auth",
		Value:  "true",
		MaxAge: 300})
}

func authEachUser(user User) {

}

func test() {
	router := mux.NewRouter()
	fs := http.StripPrefix("/static/", Logger{http.FileServer(http.Dir("./static"))})
	router.PathPrefix("/static/").Handler(fs)
	router.Handle("/auth", Logger{http.HandlerFunc(UserLoginHandler)}).Methods(http.MethodPost)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Starting server at addres: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while listening %v", err)
	}
}
