package handlers

import (
	"errors"
	"firstprogram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type counterIncrRequest struct {
	Key   string `json:"key" binding:"required"`
	Value int64  `json:"value" binding:"required"`
}

type counterIncrResponse struct {
	Value int64 `json:"value"`
}

// CounterHandler godoc
// @Summary  Инкремент значения по ключу
// @Accept   json
// @Produce  json
// @Param    request body     counterIncrRequest  true "Ключ и значение"
// @Success  200     {object} counterIncrResponse
// @Router   /counter/incr [post]
func CounterIncrHandler(counterService services.ICounterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req counterIncrRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := counterService.IncrBy(c.Request.Context(), req.Key, req.Value)
		if err != nil {
			var validErr *services.ValidationError
			if errors.As(err, &validErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, counterIncrResponse{Value: result})
	}
}
