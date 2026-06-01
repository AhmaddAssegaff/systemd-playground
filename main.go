package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request masuk dari %s... Mulai proses hello world berat / lama.", r.RemoteAddr)

		time.Sleep(10 * time.Second)

		w.Write([]byte("Hello world"))
		log.Println("Request selesai diproses dengan sukses.")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{Addr: ":" + port}

	go func() {
		log.Println("Server berjalan di port :" + port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdownSig
	log.Printf("Menerima sinyal %v. Memulai Graceful Shutdown...\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Gagal shutdown secara bersih: %v", err)
	}

	log.Println("Server berhasil dimatikan dengan aman.")
}
