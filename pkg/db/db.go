package db

import (
	"context"
	"time"

	"github.com/nicolasgomezaragon/FinancialAPI/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveToMongoDB saves financial data to MongoDB using the official driver
func SaveToMongoDB(data map[string]models.DailyData, dbName, collectionName string) error {
	// 1. Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	// 2. Get collection reference
	collection := client.Database(dbName).Collection(collectionName)

	// 3. Prepare documents for bulk insert
	var documents []interface{}
	for dateStr, dailyData := range data {
		// Parse date string to time.Time if needed
		date, err := time.Parse("2006-01-02", dateStr) // Adjust format to match your date string
		if err != nil {
			return err
		}

		documents = append(documents, bson.M{
			"date":   date,
			"open":   dailyData.Open,
			"high":   dailyData.High,
			"low":    dailyData.Low,
			"close":  dailyData.Close,
			"volume": dailyData.Volume,
		})
	}

	// 4. Insert all documents in one operation (more efficient)
	_, err = collection.InsertMany(context.TODO(), documents)
	return err
}
