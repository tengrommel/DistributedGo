package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var client *redis.Client
var store = sessions.NewCookieStore([]byte("top-s3cr3t"))
var templates *template.Template

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/",  r)
	http.ListenAndServe(":8080", nil)
}

func registerGetHandler(writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "register.html", nil)
}

func registerPostHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil{
		return
	}
	client.Set("user:" + username, hash, 0)
	http.Redirect(writer, request, "/login", 302)
}

func loginGetHandler(writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "login.html", nil)
}

func loginPostHandler(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")
	hash, err := client.Get("user:"+ username).Bytes()
	if err != nil{
		return
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil{
		return
	}
	session, _ := store.Get(request, "session")
	session.Values["username"] = username
	session.Save(request, responseWriter)
	http.Redirect(responseWriter, request, "/", 302)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok{
		http.Redirect(w, r, "/login", 302)
		return
	}
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil{
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	client.LPush("comments", comment)
	http.Redirect(w, r, "/", 302)
}