package main

import (
	"log"
	"net/http"
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

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
