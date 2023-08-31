package netxd_dal_services

import (
	"context"
	"fmt"
	"log"
	"time"
	interfaces "github.com/Sona-28/netxd_dal/netxd_dal_interfaces"
	models "github.com/Sona-28/netxd_dal/netxd_dal_models"
	"github.com/Sona-28/netxd_customer_controllers/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	ctx    context.Context
	cilent *mongo.Client
}

func InitTransaction(mclient *mongo.Client, ctx context.Context) interfaces.TransactionInterface {
	return &TransactionService{ctx, mclient}
}

func (t *TransactionService) TransferMoney(from int64, to int64, amount int64) (string, error) {
	session, err := t.cilent.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(t.ctx)
	_,err=session.WithTransaction(context.Background(),func(ctx mongo.SessionContext) (interface{}, error){
		cust:=t.cilent.Database(constants.Dbname).Collection("customer")
		trans := t.cilent.Database(constants.Dbname).Collection("transaction")
		res, err := cust.UpdateOne(ctx, bson.M{"customer_id": from}, bson.M{"$inc": bson.M{"balance": -amount}})
		if err != nil {
			return nil, err
		}
		fmt.Println(res)
		res, err = cust.UpdateOne(ctx, bson.M{"customer_id": to}, bson.M{"$inc": bson.M{"balance": amount}})
		if err != nil {
			return nil, err
		}
		res1, err := trans.InsertOne(ctx, &models.Transaction{Id:primitive.NewObjectID(), From: from, To: to, Amount: amount, TimeStamp: time.Now()})
		if err != nil {
			return nil, err
		}
		fmt.Println(res1)
		return res, nil	

		
})
	if err!=nil{
		return "",err
	}
	return "Transaction Successfull", nil
}