package dto

import (
	"assignment2/entity"
	"time"
)

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt" `
	CustomerName string           `json:"customerName" binding:"required" `
	Items        []NewItemRequest `json:"items" binding:"required"`
}

func (o *NewOrderRequest) OrderRequestToEntity() *entity.Order {
	return &entity.Order{
		CustomerName: o.CustomerName,
		OrderedAt:    o.OrderedAt,
	}
}

type NewOrderResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       NewOrderRequest `json:"data"`
}

type GetAllOrdersResponse struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Data       []OrderDataResponse `json:"data"`
}

type GetOrderByIDResponse struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       OrderDataResponse `json:"data"`
}

type DeleteOrderByIDResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type OrderDataResponse struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	CustomerName string     `json:"customerName"`
	Items        []ItemData `json:"items"`
}
