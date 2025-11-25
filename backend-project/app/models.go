package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeatherForecast struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	TemperatureF int       `json:"temperatureF"`
	Summary      string    `json:"summary"`
}

type ColorModel struct {
	Color string `json:"color" bson:"color"`
}

type TodoItem struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Done bool               `json:"done" bson:"done"`
}

type TodolistDatabaseSettings struct {
	TodoCollectionName string `json:"TodoCollectionName"`
	ConnectionString   string `json:"ConnectionString"`
	DatabaseName       string `json:"DatabaseName"`
}
