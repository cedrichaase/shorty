package database

import (
    "context"
	"log"
	"time"

    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

    "ha.si/shorty/helper"
)

var client mongo.Client
var ctx context.Context

var uri = helper.GetEnv("MONGODB_URI", "mongodb://localhost:27017")

type ShortcutDocument struct {
    Shortcut     string `json:"shortcut" bson:"shortcut" binding:"required"`
    Url          string `json:"url" bson:"url" binding:"required"`
    Created      time.Time `json:"created" bson:"created" binding:"required"`
    LastAccessed time.Time `json:"lastAccessed" bson:"lastAccessed" binding:"required"`
}

func AddShortcut(shortcut string, url string) {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    var collection = client.Database("shorty").Collection("shortcuts")
    collection.InsertOne(ctx, bson.D{
        {"shortcut", shortcut},
        {"url", url},
        {"created", time.Now()},
        {"lastAccessed", time.Now()},
    })

    defer client.Disconnect(ctx)
}

func FindUrlByShortcut(shortcut string) string {
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    var collection = client.Database("shorty").Collection("shortcuts")

    var result ShortcutDocument
    err = collection.FindOne(ctx, bson.D{{"shortcut", shortcut}}).Decode(&result)

    defer client.Disconnect(ctx)

    return result.Url
}
