package main

import (
"encoding/json"
"fmt"
"os"
"project/pkg/api"
"project/pkg/db"
"project/pkg/processing"
"project/pkg/utils"
)

type Config struct {
APIKey  string `json:"apiKey"`
MongoDB struct {
Host       string `json:"host"`
Database   string `json:"database"`
Collections struct {
RawData    string `json:"rawData"`
CleanedData string `json:"cleanedData"`
} `json:"collections"`
} `json:"mongoDB"`
}

func loadConfig(filename string) (Config, error) {
var config Config
file, err := os.Open(filename)
if err != nil {
return config, err
}
defer file.Close()

err = json.NewDecoder(file).Decode(&config)
return config, err
}

func main() {
config, err := loadConfig("config.json")
if err != nil {
fmt.Println("Error loading config: ", err)
return
}

// Read API key
apiKey := config.APIKey

// Fetch data from API
timeSeries, err := api.FetchTimeSeries(apiKey, "IBM")
if err != nil {
fmt.Println("Error fetching time series: ", err)
return
}

// Save raw data to MongoDB
err = db.SaveToMongoDB(timeSeries.TimeSeries, config.MongoDB.Database, config.MongoDB.Collections.RawData)
if err != nil {
fmt.Println("Error saving to MongoDB: ", err)
return
}

// Process data (e.g., clean and transform)
processedData := processing.CleanAndTransform(timeSeries)

// Save processed data to MongoDB
err = db.SaveToMongoDB(processedData, config.MongoDB.Database, config.MongoDB.Collections.CleanedData)
if err != nil {
fmt.Println("Error saving processed data to MongoDB: ", err)
return
}

fmt.Println("Data processing complete!")
}
