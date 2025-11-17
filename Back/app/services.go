package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ColorService reproduit la logique du service C# (lecture de BACKGROUND_COLOR).
type ColorService struct{}

func NewColorService() *ColorService {
	return &ColorService{}
}

func (s *ColorService) Get() ColorModel {
	color := os.Getenv("BACKGROUND_COLOR")
	if color == "" {
		color = "#FFFFFF"
	}
	return ColorModel{Color: color}
}

// TodoService encapsule la collection MongoDB des todos.
type TodoService struct {
	collection *mongo.Collection
}

// NewTodoService construit le client Mongo en reprenant la même logique que le C#.
func NewTodoService(ctx context.Context, settings *TodolistDatabaseSettings) (*TodoService, error) {
	uri := mongoURIFromEnv()
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("connect mongo: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo: %w", err)
	}

	db := client.Database(settings.DatabaseName)
	coll := db.Collection(settings.TodoCollectionName)

	return &TodoService{collection: coll}, nil
}

// mongoURIFromEnv reproduit exactement la logique C# :
// si DB_* sont définies on construit l'URI, sinon on utilise mongodb://root:example@mongo:27017/
func mongoURIFromEnv() string {
	addr := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	if addr != "" && port != "" && user != "" && pass != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, pass, addr, port)
	}

	return "mongodb://root:example@mongo:27017/"
}

func (s *TodoService) GetAll(ctx context.Context) ([]TodoItem, error) {
	cursor, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []TodoItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *TodoService) GetByID(ctx context.Context, id primitive.ObjectID) (*TodoItem, error) {
	var item TodoItem
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *TodoService) Create(ctx context.Context, item *TodoItem) (*TodoItem, error) {
	res, err := s.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		item.ID = oid
	}

	return item, nil
}

func (s *TodoService) Update(ctx context.Context, id primitive.ObjectID, item *TodoItem) error {
	item.ID = id
	_, err := s.collection.ReplaceOne(ctx, bson.M{"_id": id}, item)
	return err
}

func (s *TodoService) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
