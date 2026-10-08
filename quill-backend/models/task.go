package models

import (
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
)
type Task struct {
    ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID      bson.ObjectID `bson:"userId" json:"userId"`
    Title       string        `bson:"title" json:"title"`
    Description string        `bson:"description" json:"description"`
    Priority    string        `bson:"priority" json:"priority"`
    DueDate     time.Time     `bson:"dueDate" json:"dueDate"`
    Category    string        `bson:"category" json:"category"`
    Completed   bool          `bson:"completed" json:"completed"`
    CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
    UpdatedAt   time.Time     `bson:"updatedAt" json:"updatedAt"`
}