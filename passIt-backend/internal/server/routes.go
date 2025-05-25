package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	codes "passIt/internal/passit-codes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	// gin.SetMode(gin.ReleaseMode) // Set Gin to release mode
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // TODO: Add your frontend URL from env variables
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.POST("/user", s.CreateUserHandler)
	r.GET("/login", s.LoginUserHandler)
	r.GET("/jobs", s.JobsHandler)
	r.GET("/jobs/:id", s.GetJobHandler)

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

func (s *Server) JobsHandler(c *gin.Context) {
	jobsFilePath := filepath.Join("..", "passIt-ui", "src", "jobs.json")
	file, err := os.Open(jobsFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open jobs.json"})
		return
	}

	var jobs interface{}
	if err := json.NewDecoder(file).Decode(&jobs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode jobs.json"})
		return
	}
	defer file.Close()
	c.JSON(http.StatusOK, PassItResponseBody{
		Code: codes.JobsRetrievedSuccessfully,
		Data: jobs,
	})
}

type Job struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Salary      string `json:"salary"`
	Company     struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		ContactEmail string `json:"contactEmail"`
		ContactPhone string `json:"contactPhone"`
	} `json:"company"`
}

func (s *Server) GetJobHandler(c *gin.Context) {
	jobsFilePath := filepath.Join("..", "passIt-ui", "src", "jobs.json")
	file, err := os.Open(jobsFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open jobs.json"})
		return
	}

	type JobsJSON struct {
		Jobs []Job `json:"jobs"`
	}
	var jobsJSON JobsJSON
	if err := json.NewDecoder(file).Decode(&jobsJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode jobs.json"})
		return
	}
	defer file.Close()

	jobID, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, PassItResponseBody{
			Code: codes.GetJobBadRequest,
			Data: nil,
		})
		return
	}

	for _, job := range jobsJSON.Jobs {
		if jobID == job.ID {
			c.JSON(http.StatusOK, PassItResponseBody{
				Code: codes.JobsRetrievedSuccessfully,
				Data: job,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, PassItResponseBody{
		Code: codes.JobIdNotFound,
		Data: nil,
	})
}
