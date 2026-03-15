package handlers

import (
	"errors"
	"firstprogram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signRequest struct {
	Text string `json:"text" binding:"required" example:"someText"`
	Key  string `json:"key"  binding:"required" exmaple:"someKey"`
}
type signResponse struct {
	Signature string `json:"signature"`
}

// SignHandler godoc
// @Summary  Подпись HMAC-SHA512
// @Tags     sign
// @Accept   json
// @Produce  json
// @Param    request body     signRequest  true "Текст и ключ"
// @Success  200     {object} signResponse
// @Router   /sign/hmacsha512 [post]
func SignHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req signRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		signature, err := services.SignHMACSHA512(req.Text, req.Key)
		if err != nil {
			var validErr *services.ValidationError
			if errors.As(err, &validErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		resp := signResponse{
			Signature: signature,
		}
		c.JSON(http.StatusOK, resp)
	}
}
