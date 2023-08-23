package interfaces

import "Banking/models"

type CustomerFunc interface{
	CreateCustomer(customer *models.Customer)
	UpdateCustomer(customer*models.Customer)
}