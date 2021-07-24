package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

//userIdentity - Берем ID юзера из Header authorizationHeader - и записываем в контекст
func (h *Handler) userIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader) //
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	// parse token
	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)

}

// adminIdentity -  делает проверку на роль ADMIN
func (h *Handler) adminIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader) //
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	// parse token
	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	checker, err := h.services.AdminChecker(userId)
	if !checker {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}


	c.Set(userCtx, userId)

}

// getUserId -- получение ID пользователя  из контекста
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	return idInt, nil

}
