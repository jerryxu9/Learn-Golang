package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connnectionString = "mongodb+srv://jerry:jerry@cluster0.duvvmrc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "watchlist"

// important
var collection *mongo.Collection

// init method is a spectial method that runs once at the beginning of the application
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connnectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

/* DB Helpers */
// This function is lower case b/c it's a helper method that won't be exported
// Bring the Netflix struct from the model package
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	// convert the id string to the id type that mongoDB understands
	id, _ := primitive.ObjectIDFromHex(movieId)

	// MongoDB stores documents in a binary representation called BSON that allows for easy and flexible data processing
	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/#std-label-golang-bson
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie was deleted with delete count: ", deleteCount)
}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

/* Actual Controllers */
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")

	// Use the DB helper function to get all movies
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
