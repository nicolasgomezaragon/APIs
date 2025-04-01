import pymongo
import pandas as pd

# MongoDB connection
client = pymongo.MongoClient("mongodb://localhost:27017/")
db = client["financial_data"]
cleaned_collection = db["cleaned_daily_data"]

# Load cleaned data from CSV file
df = pd.read_csv('cleaned_data.csv')

# Convert DataFrame to dictionary and insert into MongoDB
cleaned_collection.insert_many(df.to_dict('records'))

print("Cleaned data loaded into MongoDB successfully!")
