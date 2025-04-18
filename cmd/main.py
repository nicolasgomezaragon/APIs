import json
from pkg import api, db

class Config:
    def __init__(self, api_key, mongo_host, mongo_database, raw_data_collection, cleaned_data_collection):
        self.api_key = api_key
        self.mongo_host = mongo_host
        self.mongo_database = mongo_database
        self.raw_data_collection = raw_data_collection
        self.cleaned_data_collection = cleaned_data_collection

def load_config(filename):
    with open(filename, 'r') as file:
        config_data = json.load(file)
        mongo_db = config_data['mongoDB']
        collections = mongo_db['collections']
        return Config(
            config_data['apiKey'],
            mongo_db['host'],
            mongo_db['database'],
            collections['rawData'],
            collections['cleanedData']
        )

def main():
    try:
        config = load_config('config.ini')
    except Exception as e:
        print(f"Error loading config: {e}")
        return

    # Read API key
    api_key = config.api_key

    # Fetch data from API
    try:
        time_series = api.fetch_time_series(api_key, 'IBM')
    except Exception as e:
        print(f"Error fetching time series: {e}")
        return

    # Save raw data to MongoDB
    try:
        db.save_to_mongo_db(time_series['TimeSeries'], config.mongo_database, config.raw_data_collection)
    except Exception as e:
        print(f"Error saving to MongoDB: {e}")
        return

    print("Data saved succesfully!")

if __name__ == "__main__":
    main()
