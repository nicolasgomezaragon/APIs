package api

import (
"encoding/json"
"fmt"
"net/http"
"project/pkg/models"
)

func FetchTimeSeries(apiKey, symbol string) (models.TimeSeries, error) {
url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)
resp, err := http.Get(url)
if err != nil {
return models.TimeSeries{}, err
}
defer resp.Body.Close()

var timeSeries models.TimeSeries
err = json.NewDecoder(resp.Body).Decode(&timeSeries)
if err != nil {
return models.TimeSeries{}, err
}

return timeSeries, nil
}
