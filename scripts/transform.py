import pymongo
import pandas as pd

# MongoDB connection
client = pymongo.MongoClient("mongodb://localhost:27017/")
db = client["financial_data"]
collection = db["daily_data"]

# Retrieve data from MongoDB
data = list(collection.find())

# Convert to DataFrame
df = pd.DataFrame(data)

# Drop the MongoDB-specific '_id' column
df.drop(columns=['_id'], inplace=True)

# Convert date column to datetime
df['date'] = pd.to_datetime(df['date'])

# Convert numeric columns to float
numeric_columns = ['open', 'high', 'low', 'close', 'volume']
for col in numeric_columns:
    df[col] = pd.to_numeric(df[col], errors='coerce')

# Handle missing values (e.g., fill with the mean or drop)
df.fillna(df.mean(), inplace=True)

# Feature engineering (e.g., calculate moving averages)
df['moving_average_5'] = df['close'].rolling(window=5).mean()
df['moving_average_10'] = df['close'].rolling(window=10).mean()

print(df.head())

# Save cleaned data to CSV file
df.to_csv('cleaned_data.csv', index=False)
