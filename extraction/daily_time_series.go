
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const symbol = "IBM"

type TimeSeries struct {
	MetaData MetaData `json:"Meta Data"`
	TimeSeries map[string]Dailydata `json:"Time Series (Daily)"`
} 

type MetaData struct {
	Information string `json: "1. Information"`
	Symbol string `json: "2. Symbol"`
	LastRefreshed string `json: "3. Last Refreshed"`
	OutputSize string `json: "4. Output Size"`
	TimeZone string `json: "5. Time Zone"` 
}

type DailyData struct {
	Open string `json:"1. open"`
	High string `json:"2. high"`
	Low string `json:"3. low"`
	Close string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func main(){
	apiKey, err := readToken(".token.txt")
	if err != nil {
		fmt.Println("Error reading token: ", err)
		return
	}
	
    url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close() // What does this line do?

	var timeSeries TimeSeries {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Meta Data: ", timeSeries.MetaData)
	for date, data := range timeSeries.TimeSeries {
		fmt.printf("Date: %s, Open: %s, High: %s, Low: %s, Close: %s, Volume: %s \n",
					date, data.Open, data.High, data.Low, data.Close, data.Volume)
	}
}

func readToken(filename string) (string, error){
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	token := strings.TrimSpace(scanner.Text())

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return token, nil
}

#VM0SM5AWTRCSIG6T