package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/z66x/graph-visualiser/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}))
	r.Post("/api/bfs", handlers.BFS)
	r.Post("/api/dfs", handlers.DFS)
	r.Post("/api/dijkstra", handlers.Dijkstra)
	
	fmt.Println("backend running on :8080")
	http.ListenAndServe(":8080", r)
}