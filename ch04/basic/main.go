package main

import (
	"net/http"
	"fmt"
)

type myHandler struct {
	greeting string
}

func (mh myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(fmt.Sprintf("%v", mh.greeting)))
}

func main() {
	//http.Handle("/", &myHandler{greeting:"Hello"})
	http.ListenAndServe(":8888", &myHandler{"Hekk"})
}
