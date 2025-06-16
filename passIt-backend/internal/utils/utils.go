package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecodeServerInput[T any](c *gin.Context, input *T) bool {
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		log.Println("JSON decode error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		log.Println("Error decoding input:", err)
		return false
	}
	return true
}
