package netxd_dal_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct{
	Id primitive.ObjectID `json:"id" bson:"id"`
	From int64 `json:"from" bson:"from"`
	To int64 `json:"to" bson:"to"`
	Amount int64 `json:"amount" bson:"amount"`
	TimeStamp time.Time `json:"timeStamp" bson:"timeStamp"`
}