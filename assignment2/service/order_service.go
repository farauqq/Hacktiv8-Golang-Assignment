package service

import (
	"assignment2/dto"
	"assignment2/entity"
	"assignment2/pkg/errs"
	"assignment2/repository/order_repository"
	"net/http"
)

type orderService struct {
	orderRepo order_repository.OrderRepository
}

type OrderService interface {
	CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr)
	GetOrderByID(orderID uint) (*dto.GetOrderByIDResponse, errs.MessageErr)
	UpdateOrderByID(orderID uint, payload dto.NewOrderRequest) (*dto.GetOrderByIDResponse, errs.MessageErr)
	DeleteOrderByID(orderID uint) (*dto.DeleteOrderByIDResponse, errs.MessageErr)
}

func NewOrderService(orderRepo order_repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, items := range payload.Items {
		item := items.ItemRequestToEntity()
		itemsPayload = append(itemsPayload, *item)
	}

	orderRequest, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create")
	}

	response := &dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data: dto.NewOrderRequest{
			OrderedAt:    orderRequest.OrderedAt,
			CustomerName: orderRequest.CustomerName,
			Items:        payload.Items,
		},
	}

	return response, nil
}

func (o *orderService) GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr) {
	orders, err := o.orderRepo.GetAllOrders()

	if err != nil {
		return nil, err
	}

	data := []dto.OrderDataResponse{}
	for _, order := range orders {
		items := []dto.ItemData{}
		for _, item := range order.Items {
			item := dto.ItemData{
				ID:          item.ID,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderID:     item.OrderID,
			}
			items = append(items, item)
		}
		order := dto.OrderDataResponse{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			CustomerName: order.CustomerName,
			Items:        items,
		}
		data = append(data, order)
	}
	response := &dto.GetAllOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       data,
	}

	return response, nil
}

func (o *orderService) GetOrderByID(orderID uint) (*dto.GetOrderByIDResponse, errs.MessageErr) {
	order, err := o.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	items := []dto.ItemData{}
	for _, eachItem := range order.Items {
		item := dto.ItemData{
			ID:          eachItem.ID,
			CreatedAt:   eachItem.CreatedAt,
			UpdatedAt:   eachItem.UpdatedAt,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderID:     eachItem.OrderID,
		}

		items = append(items, item)
	}

	response := &dto.GetOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data: dto.OrderDataResponse{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			CustomerName: order.CustomerName,
			Items:        items,
		},
	}

	return response, nil
}

func (o *orderService) UpdateOrderByID(orderID uint, payload dto.NewOrderRequest) (*dto.GetOrderByIDResponse, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, eachItem := range payload.Items {
		item := eachItem.ItemRequestToEntity()
		itemsPayload = append(itemsPayload, *item)
	}

	updatedOrder, err := o.orderRepo.UpdateOrderByID(orderID, orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	items := []dto.ItemData{}
	for _, eachItem := range updatedOrder.Items {
		item := dto.ItemData{
			ID:          eachItem.ID,
			CreatedAt:   eachItem.CreatedAt,
			UpdatedAt:   eachItem.UpdatedAt,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderID:     eachItem.ID,
		}
		items = append(items, item)
	}

	response := &dto.GetOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: dto.OrderDataResponse{
			ID:           updatedOrder.ID,
			CreatedAt:    updatedOrder.CreatedAt,
			UpdatedAt:    updatedOrder.UpdatedAt,
			CustomerName: updatedOrder.CustomerName,
			Items:        items,
		},
	}

	return response, nil
}

func (o *orderService) DeleteOrderByID(orderID uint) (*dto.DeleteOrderByIDResponse, errs.MessageErr) {
	if err := o.orderRepo.DeleteOrderByID(orderID); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    "Success delete",
	}

	return deleteResponse, nil
}
