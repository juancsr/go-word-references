package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Words struct {
    ID          primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
    Word        string              `json:"word" bson:"word"`
    Language    string              `json:"language" bson:"language"`
}
