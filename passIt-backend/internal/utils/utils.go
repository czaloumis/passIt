package utils

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
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

func InsecureHttpContext(ctx context.Context) context.Context {
	// Create an insecure HTTP client (for self-signed certs only)
	insecureHttpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 👈 disables cert check
		},
	}

	// Use the custom HTTP client in the context
	return oidc.ClientContext(ctx, insecureHttpClient)
}
