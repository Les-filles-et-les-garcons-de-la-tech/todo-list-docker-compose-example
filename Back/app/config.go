package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type appSettings struct {
	TodolistDatabaseSettings TodolistDatabaseSettings `json:"TodolistDatabaseSettings"`
}

func LoadTodolistSettings(path string) (*TodolistDatabaseSettings, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Remove UTF-8 BOM if present
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	var cfg appSettings
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}

	return &cfg.TodolistDatabaseSettings, nil
}
