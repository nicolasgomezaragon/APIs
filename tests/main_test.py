import os
import tempfile
import unittest
from pymongo import MongoClient
from pkg.utils import read_token
from pkg.db import save_to_mongo_db
from pkg.models import DailyData

class TestUtils(unittest.TestCase):
    def test_read_token(self):
        # Create a temporary token file
        with tempfile.NamedTemporaryFile(delete=False) as temp_file:
            temp_file.write(b"test_token")
            temp_file_name = temp_file.name
        
        # Test read_token function
        token, err = read_token(temp_file_name)
        self.assertIsNone(err)
        self.assertEqual(token, "test_token")
        
        # Clean up
        os.remove(temp_file_name)

class TestDB(unittest.TestCase):
    def test_save_to_mongo_db(self):
        # Mock data
        data = {
            "2025-03-31": DailyData(open="100", high="110", low="90", close="105", volume="1000")
        }
        
        # Connect to MongoDB
        client = MongoClient('localhost', 27017)
        db_name = "test_db"
        collection_name = "test_collection"
        db = client[db_name]
        collection = db[collection_name]
        
        # Clean up after test
        self.addCleanup(client.drop_database, db_name)
        
        # Test save_to_mongo_db function
        err = save_to_mongo_db(data, db_name, collection_name)
        self.assertIsNone(err)
        
        # Verify data was saved
        result = collection.find_one({"date": "2025-03-31"})
        self.assertIsNotNone(result)
        self.assertEqual(result["open"], "100")
        self.assertEqual(result["close"], "105")

if __name__ == '__main__':
    unittest.main()
