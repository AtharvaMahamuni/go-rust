package routes

import (
	"github.com/atharvamahamuni/golang/projects/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:order_item", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_item", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:order_item", controllers.UpdateOrderItem())
}
