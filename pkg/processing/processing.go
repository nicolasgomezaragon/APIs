package processing

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FinancialData represents the structure of your financial records
type FinancialData struct {
	Date    time.Time `bson:"date"`
	Open    float64   `bson:"open"`
	High    float64   `bson:"high"`
	Low     float64   `bson:"low"`
	Close   float64   `bson:"close"`
	Volume  float64   `bson:"volume"`
	MA5     float64   `bson:"moving_average_5,omitempty"`
	MA10    float64   `bson:"moving_average_10,omitempty"`
}

func CleanAndTransform() error {
	// MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("financial_data")
	collection := db.Collection("daily_data")

	// Retrieve data from MongoDB
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	var records []FinancialData
	if err = cursor.All(context.TODO(), &records); err != nil {
		return err
	}

	// Process data
	for i := range records {
		// Calculate moving averages (window of 5 and 10)
		if i >= 4 {
			records[i].MA5 = calculateMovingAverage(records, i, 5)
		}
		if i >= 9 {
			records[i].MA10 = calculateMovingAverage(records, i, 10)
		}
	}

	// Save cleaned data back to MongoDB
	cleanedCollection := db.Collection("cleaned_daily_data")
	var cleanData []interface{}
	for _, r := range records {
		cleanData = append(cleanData, r)
	}
	if _, err = cleanedCollection.InsertMany(context.TODO(), cleanData); err != nil {
		return err
	}

	// Save to CSV file
	if err = saveToCSV(records); err != nil {
		return err
	}

	return nil
}

func calculateMovingAverage(data []FinancialData, currentIndex, window int) float64 {
	var sum float64
	for i := currentIndex; i > currentIndex-window; i-- {
		sum += data[i].Close
	}
	return sum / float64(window)
}

func saveToCSV(data []FinancialData) error {
	file, err := os.Create("cleaned_data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"date", "open", "high", "low", "close", "volume", "moving_average_5", "moving_average_10"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data
	for _, record := range data {
		row := []string{
			record.Date.Format(time.RFC3339),
			strconv.FormatFloat(record.Open, 'f', -1, 64),
			strconv.FormatFloat(record.High, 'f', -1, 64),
			strconv.FormatFloat(record.Low, 'f', -1, 64),
			strconv.FormatFloat(record.Close, 'f', -1, 64),
			strconv.FormatFloat(record.Volume, 'f', -1, 64),
			strconv.FormatFloat(record.MA5, 'f', -1, 64),
			strconv.FormatFloat(record.MA10, 'f', -1, 64),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}