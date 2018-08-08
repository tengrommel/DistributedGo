package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
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
	r.HandleFunc("/test", testGetHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/",  r)
	http.ListenAndServe(":8080", nil)
}
func testGetHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "session")
	untyped, ok := session.Values["username"]
	if !ok{
		return
	}
	username, ok := untyped.(string)
	if !ok{
		return
	}
	writer.Write(([]byte(username)))
}
func loginGetHandler(writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "login.html", nil)
}

func loginPostHandler(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	session, _ := store.Get(request, "session")
	session.Values["username"] = username
	session.Save(request, responseWriter)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request)  {
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