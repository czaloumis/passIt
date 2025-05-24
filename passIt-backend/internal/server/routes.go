package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	// gin.SetMode(gin.ReleaseMode) // Set Gin to release mode
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // TODO: Add your frontend URL from env variables
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.POST("/user", s.CreateUserHandler)
	r.GET("/login", s.LoginUserHandler)

	api := r.Group("/api")

	api.Use(auth())

	api.GET("/user", s.FindUserByEmailHandler)
	api.GET("/users", s.GetAllUsersHandler)
	api.PUT("/user", s.UpdateUserByIdHandler)
	api.GET("/user/find", s.FindUserByIdHandler)

	tf := r.Group("/api/terraform")
	tf.GET("/init", s.TerraformInitHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	userInfo, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found in context"})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
