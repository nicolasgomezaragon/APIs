package models

type TimeSeries struct {
MetaData   MetaData             `json:"Meta Data"`
TimeSeries map[string]DailyData `json:"Time Series (Daily)"`
}

type MetaData struct {
Information   string `json:"1. Information"`
Symbol        string `json:"2. Symbol"`
LastRefreshed string `json:"3. Last Refreshed"`
OutputSize    string `json:"4. Output Size"`
TimeZone      string `json:"5. Time Zone"`
}

type DailyData struct {
Open   string `json:"1. open"`
High   string `json:"2. high"`
Low    string `json:"3. low"`
Close  string `json:"4. close"`
Volume string `json:"5. volume"`
}
