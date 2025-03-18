# Fetching data from 3rd party API
# To start cleaning it up

import requests

api_key=mjsbtbAPn5pu53BhqBZp2ocXNW9kT3mn # type: ignore

ticker = "BRK-A"  # Berkshire Hathaway's ticker symbol
url = f"https://financialmodelingprep.com/api/v3/quote/{ticker}?apikey={api_key}"

response = requests.get(url)
if response.status_code == 200:
    data = response.json()
    print(data)
else:
    print(f"Error: {response.status_code}")