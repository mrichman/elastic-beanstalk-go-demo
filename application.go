package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	const indexPage = "public/index.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if buf, err := ioutil.ReadAll(r.Body); err == nil {
				log.Printf("Received message: %s\n", string(buf))
			}
		} else {
			log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
			http.ServeFile(w, r, indexPage)
		}
	})

	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
