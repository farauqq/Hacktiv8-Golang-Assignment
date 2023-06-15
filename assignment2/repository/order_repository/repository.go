package order_repository

import (
	"assignment2/entity"
	"assignment2/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload *entity.Order, itemPayload []entity.Item) (*entity.Order, errs.MessageErr)
	GetAllOrders() ([]entity.Order, errs.MessageErr)
	GetOrderByID(orderID uint) (*entity.Order, errs.MessageErr)
	UpdateOrderByID(orderID uint, orderPayload *entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr)
	DeleteOrderByID(orderID uint) errs.MessageErr
}
