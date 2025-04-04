package db

import (
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"project/pkg/models"
)

func SaveToMongoDB(data map[string]models.DailyData, dbName, collectionName string) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()
	
	collection := session.DB(dbName).C(collectionName)
	
	for date, dailyData := range data {
		err = collection.Insert(bson.M{
			"date":   date,
			"open":   dailyData.Open,
			"high":   dailyData.High,
			"low":    dailyData.Low,
			"close":  dailyData.Close,
			"volume": dailyData.Volume,
		})
		if err != nil {
			return err
		}
	}
	
	return nil
}
