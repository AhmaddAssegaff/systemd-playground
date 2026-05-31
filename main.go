package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("worker is running...")
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %s", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server started on :" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
