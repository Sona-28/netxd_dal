package netxd_dal

import (
	"context"
	"fmt"
	"log"
	interfaces "netxd_dal/netxd_dal_interfaces"
	models "netxd_dal/netxd_dal_models"
	"reflect"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cust struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &Cust{ctx, collection}
}
func (c *Cust) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.D{Key: "customer_id", Value: 1}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := c.mongoCollection.Indexes().CreateMany(c.ctx, indexModel)
	if err != nil {
		return nil, err
	}
	date := time.Now()
	user.CreatedAt = date.Format("2006-01-02 15:04:05")
	user.UpdatedAt = date.Format("2006-01-02 15:04:05")
	user.IsActive = true
	_, err = c.mongoCollection.InsertOne(c.ctx, &user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Fatal("Duplicate key error")
		}
		return nil, err
	}
	result := &models.CustomerResponse{
		CustomerId: user.CustomerId,
		CreatedAt:  user.CreatedAt,
	}
	return result, nil
}

func (c *Cust) GetCustomer(id int64) (*models.Customer, error) {
	filter := bson.D{{Key: "customer_id", Value: id}}
	var customer *models.Customer
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *Cust) UpdateCustomer(val *models.UpdateCustomerRequest) (*models.CustomerResponse, error) {
	iv := bson.M{"customer_id": val.CustomerId}
	if reflect.TypeOf(val.NewValue).String() == "float64" {
		val.NewValue = int64(val.NewValue.(float64))
	}
	get,err := c.GetCustomer(val.CustomerId)
	if err!=nil {
		fmt.Println("error")
	}
	updatetime := time.Now()
	get.UpdatedAt = updatetime.Format("2006-01-02 15:04:05")
	fv := bson.M{"$set": bson.M{val.Topic: val.NewValue, "updated_at": get.UpdatedAt}}
	_, err = c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	return &models.CustomerResponse{
		CustomerId: get.CustomerId,
		CreatedAt:  get.CreatedAt,
	}, nil
}

func (c *Cust) DeleteCustomer(id int64) (*models.CustomerResponse, error) {
	del := bson.M{"customer_id": id}
	get,err := c.GetCustomer(id)
	if err!=nil {
		fmt.Println("error")
	}

	_, err = c.mongoCollection.DeleteOne(c.ctx, del)
	if err != nil {
		return nil, err
	}
	return &models.CustomerResponse{
		CustomerId: get.CustomerId,
		CreatedAt:  get.CreatedAt,
	}, nil
}
