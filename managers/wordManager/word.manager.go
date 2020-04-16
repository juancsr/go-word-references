package wordManager

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    conn "github.com/juancsr/go-word-references/db"
    "github.com/juancsr/go-word-references/models"
    "log"
    "time"
)

type ResponseFormat struct {
    Time    time.Time       `json:"time"`
    Data    interface{}     `json:"data"`
    Error   interface{}     `json:"error"`
}

func (f ResponseFormat) formatData(data interface{}) ResponseFormat {
    f.Time = time.Now()
    f.Data = data
    return f
}

func (f ResponseFormat) formatCustomError(customError interface{}) ResponseFormat {
    f.Error = customError
    return f
}

func (f ResponseFormat) noDataResponse() ResponseFormat {
    f.Error = "no data found"
    return f
}

var collectionName string = "word"

/*func CreateAll() {
    connection := conn.connect()
}*/

func GetOne(id string) ResponseFormat {
    var word *models.Words
    response := ResponseFormat{}

    client := conn.Connect()
    collection := client.Database(conn.DATABASE_NAME).Collection(collectionName)
    
    objectId, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        log.Fatal(err)
        return response.formatCustomError(err.Error())
    }

    collection.FindOne(conn.GetContext(), bson.D{{"_id", objectId}}).Decode(&word)
    if word == nil {
        return response.noDataResponse()
    }

    defer (*client).Disconnect(conn.GetContext())
    return response.formatData(word)
}

