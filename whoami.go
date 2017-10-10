package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

// DefaultPort is the port we will listen on if no config value is specified
const DefaultPort = "8080"

func respond(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		log.Printf("can't retrieve hostname: %v", err)
	}
	fmt.Fprintf(w, "Hello from %s!\n", host)
	fmt.Fprintf(w, "I'm a node running %s on CPU architecture %s", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", respond)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = DefaultPort
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Can't start webserver: ", err)
	}
}
