package netxddalinterfaces

import (
	models "netxd_dal/netxd_dal_models"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error)
	GetCustomer(customerId int64) (*models.Customer, error)
	UpdateCustomer(updateCustomerRequest *models.UpdateCustomerRequest) (*models.CustomerResponse, error)
	DeleteCustomer(customerId int64) (*models.CustomerResponse, error)
}