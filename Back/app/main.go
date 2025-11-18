package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Chargement de la config des todos depuis appsettings.json (si présent)
	settings, err := LoadTodolistSettings("appsettings.json")
	if err != nil {
		log.Printf("could not load appsettings.json, falling back to defaults: %v", err)
	}

	if settings == nil {
		// Même valeurs que dans appsettings.json de ton projet .NET
		settings = &TodolistDatabaseSettings{
			TodoCollectionName: "todo",
			DatabaseName:       "webapp",
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todoService, err := NewTodoService(ctx, settings)
	if err != nil {
		log.Fatalf("cannot create todo service: %v", err)
	}

	colorService := NewColorService()

	server := &Server{
		todoService:  todoService,
		colorService: colorService,
	}

	mux := http.NewServeMux()

	// WeatherForecast (on expose les deux variantes pour rester tolérant à la casse)
	mux.HandleFunc("/weatherforecast", server.handleWeatherForecast)
	mux.HandleFunc("/WeatherForecast", server.handleWeatherForecast)

	// Color
	mux.HandleFunc("/api/color", server.handleColorGet)

	// Todo
	mux.HandleFunc("/api/todo", server.handleTodoRoot)   // GET (liste) / POST
	mux.HandleFunc("/api/todo/", server.handleTodoByID) // GET/PUT/DELETE par id

	// CORS global "allow all" comme dans Startup.cs
	handler := withCORS(mux)

	addr := ":81"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// Middleware CORS très simple, équivalent au WithOrigins("*").AllowAnyHeader().AllowAnyMethod()
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
