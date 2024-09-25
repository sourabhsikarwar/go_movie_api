package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movies struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie_name,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
