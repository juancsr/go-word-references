package db

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
    "context"
    "log"
)

var DATABASE_NAME string = "words"

/**
* Perfomrs a very simple and default connection
*/
func Connect() *mongo.Client {
    MONGO_DB_URI := "mongodb://localhost:27017"
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_DB_URI))
    
    if err != nil {
        log.Panicln(err)
    }
    defer func() {
        if r := recover(); r != nil {
            log.Print(r)     
        }
    }()
    return client
}

func GetContext() context.Context {
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    return ctx
}

