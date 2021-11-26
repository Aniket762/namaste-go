package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Aniket762/namaste-go/dbapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// update 1 record
func UpdateOneMovie(movieId string)  {
	// converting string into valid mongo id
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil{
		log.Fatal(err)
	}

	// creating filter to find the exact movie
	// everything in mongodb is bson
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(),filter,update)

	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("modified count: ",result.ModifiedCount)

}

// delete 1 record
func DeleteOneMovie(movieId string)  {
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount , _ := collection.DeleteOne(context.Background(),filter)
	fmt.Println("Movie got delete with delete count: ",deleteCount)
}

// delete all records
func DeleteAllMovies() int64 {
	// empty params implies all
	deleteResult , err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("The movies which were deleted: ",deleteResult)
	fmt.Println("The number of movies deleted: ",deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

// get all movies from db
func GetAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(),bson.M{})
	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M 
	
	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// controllers
func GetAllMoviesController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
 	allMovies := GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatch(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])	
}

func DeleteOneMovieController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovieController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	
	count := DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
