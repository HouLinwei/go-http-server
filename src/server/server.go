package main

import (
	"log"
	"os"
	"net/http"
)

var logger *log.Logger

func main() {

	// Create a logger
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(" --- start my server --- ")

	// Register a handler
	http.HandleFunc("/", IndexHandler)

	// Start my server
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(`hello world!`))
}