package handler

import (
	"github.com/nekruz08/online-store/pkg/serivce"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	i := router.Group("/items")
	i.GET("/", h.getAllItems)
	i.POST("/category", h.getItemsById)

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}

	adminApi := router.Group("/admin/api", h.adminIdentity)
	{
		items := adminApi.Group("/items")
		items.POST("/", h.createItem)
		items.DELETE("/", h.deleteItemByID)
		items.PUT("/", h.updateItemByID)

		orders := adminApi.Group("/order")
		{
			orders.POST("/confirm", h.confirmOrder)
			orders.GET("/delivered", h.getDelivered)
			orders.GET("/active", h.getActiveOrders)
		}
	}



	// USERS PART
	api := router.Group("/api", h.userIdentity)
	{
		carts := api.Group("/carts")
		{
			carts.POST("/", h.addCarts)
			carts.GET("/", h.getItemsFromCart)
			carts.PUT("/", h.updateCart)

		}
		api.GET("/order", h.getOrderByID)
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
		}
	}

	return router
}
