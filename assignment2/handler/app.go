package handler

import (
	"assignment2/database"
	"assignment2/dto"
	"assignment2/pkg/errs"
	"assignment2/repository/item_repository/item_postgres"
	"assignment2/repository/order_repository/order_postgres"

	"assignment2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := database.GetDataBaseInstance()

	itemRepo := item_postgres.NewItemPG(db)
	orderRepo := order_postgres.NewOrderPG(db, itemRepo)
	orderService := service.NewOrderService(orderRepo)

	r := gin.Default()

	r.POST("/orders", func(ctx *gin.Context) {
		var request dto.NewOrderRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			newError := errs.NewUnprocessableEntity(err.Error())
			ctx.JSON(newError.StatusCode(), newError)
			return
		}

		order, err := orderService.CreateOrder(request)
		if err != nil {
			ctx.JSON(err.StatusCode(), err)
			return
		}

		ctx.JSON(order.StatusCode, order)
	})

	r.GET("/orders", func(ctx *gin.Context) {
		orders, err := orderService.GetAllOrders()
		if err != nil {
			ctx.JSON(err.StatusCode(), err)
			return
		}

		ctx.JSON(orders.StatusCode, orders)
	})

	r.GET("/orders/:orderID", func(ctx *gin.Context) {
		orderID := ctx.Param("orderID")
		orderIDInt, err := strconv.Atoi(orderID)
		if err != nil {
			newError := errs.NewBadRequest("orderID should be an unsigned integer")
			ctx.JSON(newError.StatusCode(), newError)
			return
		}

		order, errOrder := orderService.GetOrderByID(uint(orderIDInt))
		if errOrder != nil {
			ctx.JSON(errOrder.StatusCode(), errOrder)
			return
		}
		ctx.JSON(order.StatusCode, order)
	})

	r.PATCH("/orders/:orderID", func(ctx *gin.Context) {
		orderID := ctx.Param("orderID")
		orderIDInt, err := strconv.Atoi(orderID)
		if err != nil {
			newError := errs.NewBadRequest("orderID should be an unsigned integer")
			ctx.JSON(newError.StatusCode(), newError)
			return
		}

		var request dto.NewOrderRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			newError := errs.NewUnprocessableEntity(err.Error())
			ctx.JSON(newError.StatusCode(), newError)
			return
		}

		updatedOrder, errOrder := orderService.UpdateOrderByID(uint(orderIDInt), request)
		if errOrder != nil {
			ctx.JSON(errOrder.StatusCode(), errOrder)
			return
		}

		ctx.JSON(updatedOrder.StatusCode, updatedOrder)
	})

	r.DELETE("/orders/:orderID", func(ctx *gin.Context) {
		orderID := ctx.Param("orderID")
		orderIDInt, err := strconv.Atoi(orderID)
		if err != nil {
			newError := errs.NewBadRequest("orderID should be an unsigned integer")
			ctx.JSON(newError.StatusCode(), newError)
			return
		}

		deleteOrder, errOrder := orderService.DeleteOrderByID(uint(orderIDInt))
		if errOrder != nil {
			ctx.JSON(errOrder.StatusCode(), errOrder)
			return
		}
		ctx.JSON(deleteOrder.StatusCode, deleteOrder)
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}
