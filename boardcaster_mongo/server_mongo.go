package boardcaster_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

//docker run --name some-mongo -p 27017:27017 -v my-vol:/data/db -it mongo
//docker run -it mongo mongo --host 172.17.0.2

const (
	MONGO_URL = "mongodb://127.0.0.1:27017"
)

//Collection is a handle to a MongoDB collection.
//It is safe for concurrent use by multiple goroutines.
var mcollection *mongo.Collection
var ucollection *mongo.Collection

func MongoInit() {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mcollection = client.Database("im").Collection("message")
	ucollection = client.Database("im").Collection("online_user")
}

func GetAllUser() []User {
	results := make([]User, 0)
	cur, err := ucollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			Name string
			ID   primitive.ObjectID `bson:"_id"`
			Addr string
		}{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			continue
		}
		user := User{Name: result.Name, Id: result.ID}
		results = append(results, user)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func GetMessage(userId string) []Message {
	results := make([]Message, 0)
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"user_Id", objId}}
	cur, err := mcollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			Content string
			Time    time.Time
		}{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			continue
		}
		mes := Message{Content: result.Content, Time: result.Time}
		results = append(results, mes)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func InsertUser(userP *User) {
	_, err := ucollection.InsertOne(context.Background(), bson.M{
		"_id":  userP.Id,
		"name": userP.Name,
		"addr": userP.Addr.String()})
	if err != nil {
		log.Fatal(err)
		log.Fatal("inser user to mongo failed")
	}
	fmt.Println("new user:", userP.Id)
}

func DeleteUser(userP *User) {
	res, err := ucollection.DeleteOne(context.TODO(), bson.D{{"name", userP.Name}})
	if err != nil {
		log.Fatal(err)
		log.Fatal("delete user from mongo failed")
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

func InsertMessage(mes Message) {
	res, err := mcollection.InsertOne(context.Background(), bson.D{
		bson.E{Key: "content", Value: mes.Content},
		bson.E{"time", mes.Time},
		bson.E{"user_Id", mes.User.Id}})
	if err != nil {
		log.Fatal("insert message to mongo failed")
	}
	id := res.InsertedID
	fmt.Println("new message:", id)
}
