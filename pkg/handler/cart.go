package handler

import (
	"github.com/nekruz08/online-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemInput struct {
	ItemId int `json:"item_id"`
	Count int `json:"count"`
}



func (h *Handler) addCarts(c *gin.Context)  {
	var input ItemInput
	id, err := getUserId(c)
	if err != nil {
		return
	}
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.CreateCarts(id,input.Count, input.ItemId)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
		}
	c.JSON(200, gin.H{
		"Message": "Товар успешно добавлен",
	})

}

type GetItemsFromCarts struct {
	Data []models.UsersCarts `json:"data"`
}


// getItemsFromCart  - берет товары из корзины по ID
func (h *Handler) getItemsFromCart(c *gin.Context)  {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	cartItems, total, err := h.services.GetCartById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"GetItemsFromCarts":GetItemsFromCarts{
			Data: cartItems,
		},
		"total": total,
	})



}

func (h *Handler) updateCart(c *gin.Context)  {
	var input ItemInput
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	// calls method
	err = h.services.UpdateCartByID(userId, input.ItemId, input.Count)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "корзина успешно обновлена",
	})


}