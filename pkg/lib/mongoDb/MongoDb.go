package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	client *mongo.Client
	credential = options.Credential{
		Username: "admin",
		Password: "admin",
	}
	databaseName = "configuration"
	groupCollection *mongo.Collection
)


var clientOpts = options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

func connect() {
	if client == nil {
		log.Println("creating client for mongo")
		var err error
		if client, err = mongo.Connect(context.TODO(), clientOpts); err != nil {
			panic(err)
		}
		if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
			log.Panic(err)
		}
		databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
		if err != nil {
			log.Println(err)
		}
		log.Println(databases)
		log.Println("connection with mongo established")
	}
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func groupOperation() *mongo.Collection{
	connect()
	if groupCollection == nil {
	groupCollection =	client.Database(databaseName).Collection("Group")
	}
	return groupCollection
}

func Find(key string) {
	dfCtx, cancel := defaultContext()
	defer cancel()
	find, err := groupCollection.Find(dfCtx, "")
	if err != nil {
		return
	}

}
