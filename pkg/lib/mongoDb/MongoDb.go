package mongoDb

import (
	"context"
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	client     *mongo.Client
	credential = options.Credential{
		Username: "admin",
		Password: "admin",
	}
	databaseName    = "configuration"
	groupCollection *mongo.Collection
)

var clientOpts = options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

func connect() {
	if client == nil {
		log.Println("creating client for mongo")
		var err error
		if client, err = mongo.Connect(context.TODO(), clientOpts); err != nil {
			log.Panic(err)
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

func groupOperation() *mongo.Collection {
	connect()
	if groupCollection == nil {
		groupCollection = client.Database(databaseName).Collection("group")
	}
	return groupCollection
}

func FindGroupById(id int, model interface{}) error {
	m := bson.M{"_id": id}
	ctx, cancel := defaultContext()
	defer cancel()
	single := groupOperation().FindOne(ctx, m)
	return single.Decode(model)
}

func CreateGroup(gp model.Group) error {
	ctx, cancel := defaultContext()
	filter := bson.D{{"_id", gp.Id}}
	update := bson.M{"$set": gp}
	opt := options.Update().SetUpsert(true)
	defer cancel()
	_, err := groupOperation().UpdateOne(ctx, filter, update, opt)
	return err
}

func DeleteGroupById(id int) error {
	ctx, cancel := defaultContext()
	defer cancel()
	if _, err := groupOperation().DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}
