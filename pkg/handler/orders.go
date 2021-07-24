package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type orderInput struct {
	Phone string `json:"phone"`
	Address string `json:"address"`
}

type CartID struct {
	Id int `json:"id"`
}
func (h *Handler) createOrder(c *gin.Context)  {

	var input orderInput
	id, err := getUserId(c)
	if err != nil {
		return
	}
	//call service method

	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.CreateOrder(input.Phone, input.Address, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Заказ Оформлен, ждите с вами свяжется курьер",
	})

}

func (h *Handler) getActiveOrders(c *gin.Context)  {

	orders, err := h.services.GetActiveOrders()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Orders": orders,
	})
}

func (h *Handler) getDelivered(c *gin.Context)  {
	orders, err := h.services.GetDeliveredOrders()
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Orders": orders,
	})
}

func (h *Handler) getOrderByID(c *gin.Context)  {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.GetOrderByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (h *Handler) confirmOrder(c *gin.Context)  {
	var cartID CartID
	err := c.BindJSON(&cartID)
	if err != nil {
		return
	}
	err = h.services.ConfirmOrder(cartID.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message:": "Клиент принял товар",
	})

}