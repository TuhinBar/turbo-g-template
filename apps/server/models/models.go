package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    Done        bool               `json:"done" bson:"done"`
    CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}