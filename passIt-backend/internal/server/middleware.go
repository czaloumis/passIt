package server

import (
	"crypto/tls"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

// TODO: Use the keycloack client from the keycloak package
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the request header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Validate the token using Keycloak
		client := gocloak.NewClient("https://localhost:8443")
		restyClient := client.RestyClient()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		userInfo, err := client.GetUserInfo(c.Request.Context(), token, "passit")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		// Store user info in context for later use
		c.Set("user", userInfo)
		c.Next()
	}
}
