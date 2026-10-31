package main

import (
	"log"
	"net/http"
	"time"

	"watchtogether/internal/features/rooms/repository/memory"
	"watchtogether/internal/features/rooms/service"
	roomshttp "watchtogether/internal/features/rooms/transport/http"
)

func main() {
	mux := http.NewServeMux()
	rooms := service.New(memory.New())
	roomshttp.New(rooms).Register(mux)
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })
	server := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Println("Watchroom API listening on :8080")
	log.Fatal(server.ListenAndServe())
}
