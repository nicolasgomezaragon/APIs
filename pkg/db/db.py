from pymongo import MongoClient
from pkg.db import models

def save_to_mongo_db(data, db_name, collection_name):
    client = MongoClient('localhost', 27017)
    db = client[db_name]
    collection = db[collection_name]
    
    for date, daily_data in data.items():
        document = {
            "date": date,
            "open": daily_data.open,
            "high": daily_data.high,
            "low": daily_data.low,
            "close": daily_data.close,
            "volume": daily_data.volume,
        }
        collection.insert_one(document)
    
    client.close()
