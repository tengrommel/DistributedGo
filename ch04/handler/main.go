package main

import (
	"net/http"
	"os"
	"log"
	"io"
	"strings"
)

// 简单的文件服务器
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		f, err := os.Open("public" + request.URL.Path)
		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		var contentType string
		switch  {
		case strings.HasSuffix(request.URL.Path,"css"):
			contentType = "text/css"
		case strings.HasPrefix(request.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(request.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"
		}
		writer.Header().Add("Content-Type", contentType)
		io.Copy(writer, f)
	})
	http.ListenAndServe(":8008", nil)
}
