package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Aniket762/namaste-go/dbapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString = ""
const dbName = "namastenetflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with db
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// mongodb controllers

// insert 1 record
func InsertOneMovie(movie model.Netflix)  {
	// context.Background() is like this considering the cluster as excecution context
	inserted, err := collection.InsertOne(context.Background(),movie)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id ", inserted.InsertedID)
}
