package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Model Definitions
type Timeseries struct {
	Date   time.Time `bson:"date"`
	Open   float64   `bson:"open"`
	High   float64   `bson:"high"`
	Low    float64   `bson:"low"`
	Close  float64   `bson:"close"`
	Volume float64   `bson:"volume"`
}

// SaveProcessedData saves the processed data into the specified MongoDB collection.
func (s *FinancialService) SaveProcessedData(ctx context.Context, database, collection string, data []ProcessedData) error {
	coll := s.db.Database(database).Collection(collection)
	var docs []interface{}
	for _, d := range data {
		docs = append(docs, d)
	}
	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		return fmt.Errorf("failed to insert documents: %w", err)
	}
	return nil
}

type ProcessedData struct {
	Date            time.Time `bson:"date"`
	Close           float64   `bson:"close"`
	MovingAverage5  float64   `bson:"moving_average_5"`
	MovingAverage10 float64   `bson:"moving_average_10"`
}

// Service Layer
type FinancialService struct {
	db *mongo.Client
}

func NewFinancialService(mongoURI string) (*FinancialService, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	return &FinancialService{db: client}, nil
}

// Data Processing
func (s *FinancialService) CleanAndTransform(timeseries []Timeseries) ([]ProcessedData, error) {
	var processed []ProcessedData

	for i, data := range timeseries {
		pd := ProcessedData{
			Date:  data.Date,
			Close: data.Close,
		}

		// Calculate moving averages
		if i >= 4 {
			pd.MovingAverage5 = calculateMovingAverage(timeseries, i, 5)
		}
		if i >= 9 {
			pd.MovingAverage10 = calculateMovingAverage(timeseries, i, 10)
		}

		processed = append(processed, pd)
	}

	return processed, nil
}

func calculateMovingAverage(data []Timeseries, currentIndex, window int) float64 {
	var sum float64
	for i := currentIndex; i > currentIndex-window; i-- {
		sum += data[i].Close
	}
	return sum / float64(window)
}

// API Client
func FetchTimeseries(apiKey, symbol string) ([]Timeseries, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Parse response into Timeseries slice
	// (Implementation depends on AlphaVantage's response structure)
	// This is a placeholder - adjust according to actual API response
	var timeseries []Timeseries
	return timeseries, nil
}

// Main Workflow
func main() {
	// Initialize services
	service, err := NewFinancialService("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Service initialization failed: %v", err)
	}

	// Fetch data
	timeSeries, err := FetchTimeseries("YOUR_API_KEY", "AAPL")
	if err != nil {
		log.Fatalf("Failed to fetch timeseries: %v", err)
	}

	// Process data
	processedData, err := service.CleanAndTransform(timeSeries)
	if err != nil {
		log.Fatalf("Data processing failed: %v", err)
	}

	// Store results
	if err := service.SaveProcessedData(context.Background(), "financial_db", "processed_data", processedData); err != nil {
		log.Fatalf("Failed to save processed data: %v", err)
	}
}