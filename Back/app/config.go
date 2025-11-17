package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type appSettings struct {
	TodolistDatabaseSettings TodolistDatabaseSettings `json:"TodolistDatabaseSettings"`
}

// LoadTodolistSettings lit appsettings.json et retourne les paramètres pour la TodoList.
func LoadTodolistSettings(path string) (*TodolistDatabaseSettings, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg appSettings
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}

	return &cfg.TodolistDatabaseSettings, nil
}
