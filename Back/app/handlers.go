package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	todoService  *TodoService
	colorService *ColorService
}

// Helper générique d'écriture JSON.
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("writeJSON error: %v", err)
		}
	}
}

// ---------------------- WeatherForecast ----------------------

var summaries = []string{
	"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching",
}

func (s *Server) handleWeatherForecast(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	forecasts := make([]WeatherForecast, 0, 5)

	for i := 1; i <= 5; i++ {
		tempC := rng.Intn(75) - 20 // -20 to 54
		tempF := 32 + int(float64(tempC)/0.5556)

		forecasts = append(forecasts, WeatherForecast{
			Date:         time.Now().AddDate(0, 0, i),
			TemperatureC: tempC,
			TemperatureF: tempF,
			Summary:      summaries[rng.Intn(len(summaries))],
		})
	}

	writeJSON(w, http.StatusOK, forecasts)
}

// ---------------------- Color ----------------------

func (s *Server) handleColorGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	color := s.colorService.Get()
	writeJSON(w, http.StatusOK, color)
}

// ---------------------- Todo ----------------------

// handleTodoRoot gère /api/todo pour GET (liste) et POST (création).
func (s *Server) handleTodoRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleTodoList(w, r)
	case http.MethodPost:
		s.handleTodoCreate(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// handleTodoByID gère /api/todo/{id} pour GET/PUT/DELETE.
func (s *Server) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	idPart := strings.TrimPrefix(r.URL.Path, "/api/todo/")
	if idPart == "" || strings.Contains(idPart, "/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Dans le C#, on a la contrainte {id:length(24)} => ObjectId Mongo
	if len(idPart) != 24 {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	oid, err := primitive.ObjectIDFromHex(idPart)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.handleTodoGet(w, r, oid)
	case http.MethodPut:
		s.handleTodoUpdate(w, r, oid)
	case http.MethodDelete:
		s.handleTodoDelete(w, r, oid)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleTodoList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	items, err := s.todoService.GetAll(ctx)
	if err != nil {
		log.Printf("GetAll error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, items)
}

func (s *Server) handleTodoGet(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	item, err := s.todoService.GetByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.NotFound(w, r)
			return
		}
		log.Printf("GetByID error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, item)
}

func (s *Server) handleTodoCreate(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var item TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// On laisse MongoDB générer l'ID si non fourni.
	item.ID = primitive.NilObjectID

	created, err := s.todoService.Create(ctx, &item)
	if err != nil {
		log.Printf("Create error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Équivalent du CreatedAtRoute("GetTodo", new { id = todo.Id })
	w.Header().Set("Location", "/api/todo/"+created.ID.Hex())
	writeJSON(w, http.StatusCreated, created)
}

func (s *Server) handleTodoUpdate(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var item TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// On vérifie d'abord que le todo existe, comme dans le contrôleur C#.
	_, err := s.todoService.GetByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.NotFound(w, r)
			return
		}
		log.Printf("GetByID before update error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := s.todoService.Update(ctx, id, &item); err != nil {
		log.Printf("Update error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleTodoDelete(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err := s.todoService.GetByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.NotFound(w, r)
			return
		}
		log.Printf("GetByID before delete error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := s.todoService.Delete(ctx, id); err != nil {
		log.Printf("Delete error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
