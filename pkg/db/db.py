from pymongo import MongoClient

def save_to_mongo_db(data, db_name: str, collection_name: str):
    client = MongoClient('localhost', 27017)
    db = client[db_name]
    collection = db[collection_name]
    
    documents = []

    for date, daily_data in data.items():
        try:
            document = {
                "date": date,
                "open": float(daily_data.open),
                "high": float(daily_data.high),
                "low": float(daily_data.low),
                "close": float(daily_data.close),
                "volume": int(daily_data.volume),
            }
            documents.append(document)
        except Exception as e:
            print(f"[!] Skipping {date} due to error: {e}")
    
    if documents:
        collection.insert_many(documents)  # batch insert for performance
    
    client.close()
