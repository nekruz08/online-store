package handler

import (
	"github.com/nekruz08/online-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


//singUp - регистрация пользоваля
func (h *Handler) singUp(c *gin.Context)  {
	var input  models.User

	if err :=  c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return

	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

type singInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

}

// singIn - авторизация
func (h *Handler) singIn (c *gin.Context)  {
	var input  singInInput

	if err :=  c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}