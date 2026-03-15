package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"firstprogram/services"
)

type createUserRequest struct {
	Name string `json:"name" binding:"required" example:"Alex"`
	Age  int    `json:"age"  binding:"required,gt=0" example:"21"`
}

type createUserResponse struct {
	ID int64 `json:"id" example:"1"`
}

// UsersHandler godoc
// @Summary  Создать пользователя
// @Accept   json
// @Produce  json
// @Param    request body     createUserRequest  true "Имя и возраст"
// @Success  200     {object} createUserResponse
// @Router   /users/create [post]
func CreateUserHandler(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := userService.CreateUser(c.Request.Context(), req.Name, req.Age)
		if err != nil {
			var validErr *services.ValidationError
			if errors.As(err, &validErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, createUserResponse{ID: id})
	}
}
