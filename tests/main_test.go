package main

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nicolasgomezaragon/FinancialAPI/pkg/models"
	"github.com/nicolasgomezaragon/FinancialAPI/pkg/utils"
	"github.com/nicolasgomezaragon/FinancialAPI/pkg/db"
)

func TestReadToken(t *testing.T) {
	// Create a temporary token file
	file, err := os.CreateTemp("", "token.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name())

	// Write a test token to the file
	token := "test_token"
	if _, err := file.WriteString(token); err != nil {
		t.Fatalf("Failed to write token: %s", err)
	}
	file.Close()

	// Test readToken function
	readToken, err := utils.ReadToken(file.Name())
	if err != nil {
		t.Fatalf("Failed to read token: %s", err)
	}

	if readToken != token {
		t.Errorf("Expected %s, got %s", token, readToken)
	}
}

func TestSaveToMongoDB(t *testing.T) {
	// Mock data
	data := map[string]models.DailyData{
		"2025-03-31": {Open: "100", High: "110", Low: "90", Close: "105", Volume: "1000"},
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %s", err)
	}
	defer client.Disconnect(context.TODO())

	// Use a test database and collection
	dbName := "test_db"
	collectionName := "test_collection"
	collection := client.Database(dbName).Collection(collectionName)

	// Clean up after test
	defer client.Database(dbName).Drop(context.TODO())

	// Test saveToMongoDB function
	err = db.SaveToMongoDB(data, dbName, collectionName)
	if err != nil {
		t.Fatalf("Failed to save to MongoDB: %s", err)
	}

	// Verify data was saved
	var result struct {
		Open  string `bson:"open"`
		Close string `bson:"close"`
	}
	err = collection.FindOne(context.TODO(), bson.M{"date": "2025-03-31"}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find data: %s", err)
	}

	if result.Open != "100" || result.Close != "105" {
		t.Errorf("Expected Open: 100, Close: 105, got Open: %s, Close: %s", result.Open, result.Close)
	}
}
