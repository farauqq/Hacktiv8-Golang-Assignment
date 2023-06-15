package order_postgres

import (
	"assignment2/entity"
	"assignment2/pkg/errs"
	"assignment2/repository/item_repository"
	"assignment2/repository/order_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type orderPg struct {
	db       *gorm.DB
	itemRepo item_repository.ItemRepository
}

func NewOrderPG(database *gorm.DB, itemRepo item_repository.ItemRepository) order_repository.OrderRepository {
	return &orderPg{
		db:       database,
		itemRepo: itemRepo,
	}
}

func (o *orderPg) CreateOrder(orderPayload *entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Printf("Error: %v\n", err.Error())
		return nil, errs.NewInternalServerError("Failed to begin transaction")
	}

	if err := tx.Create(orderPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewBadRequest(fmt.Sprintf("Failed to create new order. %v", err.Error()))
	}

	for _, item := range itemsPayload {
		if err := tx.Model(orderPayload).Association("Items").Append(&item); err != nil {
			tx.Rollback()
			return nil, errs.NewBadRequest(fmt.Sprintf("Failed to create new item. %v", err.Error()))
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Failed to commit transaction")
	}

	return orderPayload, nil
}

func (o *orderPg) GetAllOrders() ([]entity.Order, errs.MessageErr) {
	var orders []entity.Order

	if err := o.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all orders")
	}

	return orders, nil
}

func (o *orderPg) GetOrderByID(orderID uint) (*entity.Order, errs.MessageErr) {
	var order entity.Order

	if err := o.db.Preload("Items").First(&order, orderID).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Order with id %d is not found", orderID))
	}

	return &order, nil
}

func (o *orderPg) UpdateOrderByID(orderID uint, orderPayload *entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Printf("Error: %v\n", err.Error())
		return nil, errs.NewInternalServerError("Failed to begin transaction")
	}

	if err := tx.Model(order).Updates(orderPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewBadRequest(fmt.Sprintf("Order with id %d failed to update. %v", orderID, err.Error()))
	}

	order.Items = []entity.Item{}
	for _, item := range itemsPayload {
		updatedItem, err := o.itemRepo.UpdateItemByItemCode(item.ItemCode, &item, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		order.Items = append(order.Items, *updatedItem)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to commit transaction")
	}

	return order, nil
}

func (o *orderPg) DeleteOrderByID(orderID uint) errs.MessageErr {
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	if err := o.db.Delete(order).Error; err != nil {
		return errs.NewInternalServerError("Failed to delete order")
	}

	return nil
}
