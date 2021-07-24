package handler

import (
	"fmt"
	"github.com/nekruz08/online-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createItem(c *gin.Context)  {
	var input models.Item
	fmt.Println("вызвался обрабочтик")
	err:= c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "товар успешно добавлен",
	})


}

type IdCategories struct {
	Id int `json:"id"`
}
type IdItem struct {
	Id int `json:"id"`
}


func (h *Handler) getItemsById(c *gin.Context)  {
	var id IdCategories

	err := c.BindJSON(&id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.GetItemsByIdCategory(id.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})

}

func (h *Handler) getAllItems(c *gin.Context)  {

	items, err := h.services.GetAllItems()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Items": items,
	})

}

func (h *Handler) deleteItemByID(c *gin.Context)  {
	var id IdItem

	err := c.BindJSON(&id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// call method
	IdItem, err := h.services.RemoveItemByID(id.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message:":  "Товар успешно удален",
		"ID": IdItem,
	})
}

func (h *Handler) updateItemByID(c *gin.Context)  {
	var item *models.Item

	err := c.BindJSON(&item)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// call method
	newItem, err := h.services.UpdateItemByID(item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message:":  "Товар успешно обновлен",
		"Item": newItem,
	})

}