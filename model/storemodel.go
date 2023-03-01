package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemName       string             `json:"itemname,omitempty"`
	Price          int                `json:"price,omitempty"`
	StockAvailable int                `json:"stockavailable,omitempty"`
}
