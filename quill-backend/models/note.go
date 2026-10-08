package models

import (
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
)
type Note struct {
    ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    bson.ObjectID `bson:"userId" json:"userId"`
    Title     string        `bson:"title" json:"title"`
    Content   string        `bson:"content" json:"content"`
    CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
    UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

