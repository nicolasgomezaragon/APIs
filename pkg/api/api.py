import requests
import configparser

from pkg.models import models

config = configparser.ConfigParser()
config.read('config.ini')
api_key = config['API']['api_key']
symbol = config['API']['symbol']


def fetch_time_series(api_key, symbol):
    url = f"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol={symbol}&apikey={api_key}"
    response = requests.get(url)
    if response.status_code != 200:
        return {}, Exception(f"Error fetching data: {response.status_code}")
    
    time_series = response.json()
    return models.TimeSeries(**time_series), None
