package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/shubhamgangwar/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ShubhamGangwar@cluster0.byb868a.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

// most important
var collection *mongo.Collection

// connect with mongodb
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("err")
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}


//mongo helpers-file

//insert 1 record

func insertOneMovie(movie model.Netflix){
   inserted,err := collection.InsertOne(context.Background(),movie)

   if err!=nil {
	log.Fatal(err)
   }
   fmt.Println("inserted one movie in db with id:",inserted.InsertedID)
}

//update 1 record
func updateOneMovie(movieID string){
    id,_:=primitive.ObjectIDFromHex(movieID)
	filter :=bson.M{"_id":id}
	update := bson.M{"$Set":bson.M("watched":true)}
	result,err :=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}