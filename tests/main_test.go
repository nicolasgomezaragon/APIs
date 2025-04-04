package main

import (
"os"
"testing"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
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
file.WriteString(token)
file.Close()

// Test readToken function
readToken, err := readToken(file.Name())
if err != nil {
t.Fatalf("Failed to read token: %s", err)
}

if readToken != token {
t.Errorf("Expected %s, got %s", token, readToken)
}
}

func TestSaveToMongoDB(t *testing.T) {
// Mock data
data := map[string]DailyData{
"2025-03-31": {Open: "100", High: "110", Low: "90", Close: "105", Volume: "1000"},
}

// Connect to MongoDB
session, err := mgo.Dial("localhost")
if err != nil {
t.Fatalf("Failed to connect to MongoDB: %s", err)
}
defer session.Close()

// Use a test database and collection
dbName := "test_db"
collectionName := "test_collection"
collection := session.DB(dbName).C(collectionName)

// Clean up after test
defer session.DB(dbName).DropDatabase()

// Test saveToMongoDB function
err = saveToMongoDB(data, dbName, collectionName)
if err != nil {
t.Fatalf("Failed to save to MongoDB: %s", err)
}

// Verify data was saved
var result DailyData
err = collection.Find(bson.M{"date": "2025-03-31"}).One(&result)
if err != nil {
t.Fatalf("Failed to find data: %s", err)
}

if result.Open != "100" || result.Close != "105" {
t.Errorf("Expected Open: 100, Close: 105, got Open: %s, Close: %s", result.Open, result.Close)
}
}
