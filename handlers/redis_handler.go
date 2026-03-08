package handlers

import (
	"errors"
	"firstprogram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type redisIncrRequest struct {
	Key   string `json:"key" binding:"required"`
	Value int64  `json:"value" binding:"required"`
}

type redisIncrResponse struct {
	Value int64 `json:"value"`
}

// RedisIncrHandler godoc
// @Summary  Инкремент значения в Redis
// @Tags     redis
// @Accept   json
// @Produce  json
// @Param    request body     redisIncrRequest  true "Ключ и значение"
// @Success  200     {object} redisIncrResponse
// @Router   /redis/incr [post]
func RedisIncrHandler(redisSvc *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req redisIncrRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := redisSvc.IncrBy(c.Request.Context(), req.Key, req.Value)
		if err != nil {
			var validErr *services.ValidationError
			if errors.As(err, &validErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, redisIncrResponse{Value: result})
	}
}
