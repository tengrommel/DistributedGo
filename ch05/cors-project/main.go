package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request)  {
	response.Write([]byte("CORS"))
}


func main() {
	router := mux.NewRouter()
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	http.ListenAndServe(":12345", handlers.CORS(headers, methods, origins)(router))

}
