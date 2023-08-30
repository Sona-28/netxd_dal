package netxd_dal_models

type Customer struct {
	CustomerId int64 `json:"customer_id" bson:"customer_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName string `json:"last_name" bson:"last_name"`
	BankId int64 `json:"bank_id" bson:"bank_id"`
	Balance int64 `json:"balance" bson:"balance"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	IsActive bool `json:"is_active" bson:"is_active"`
}

type CustomerResponse struct {
	CustomerId int64 `json:"customer_id" bson:"customer_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}

type UpdateCustomerRequest struct {
	CustomerId int64 `json:"customer_id"`
	Topic string `json:"topic"`
	NewValue interface{} `json:"new_value"`
}
