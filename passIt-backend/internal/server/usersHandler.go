package server

import (
	"encoding/json"
	"log"
	"net/http"
	"passIt/internal/models"
	codes "passIt/internal/passit-codes"
	"passIt/keyclock"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type LoginUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FindUserByIdRequestBody struct {
	ID uuid.UUID `json:"id"`
}

type FindUserByEmailRequestBody struct {
	Email string `json:"email"`
}

type CreateUserRequestBody struct {
	User     models.User `json:"user"`
	Password string      `json:"password"`
}

type CreateUserReturnBody struct {
	User            models.User `json:"user"`
	KeycloackUserID string      `json:"keycloak_user_id"`
}

func (s *Server) LoginUserHandler(c *gin.Context) {
	var input LoginUserRequestBody
	var keycloackClient keyclock.KeycloakClient
	k := keycloackClient.NewKeycloakClient()
	json.NewDecoder(c.Request.Body).Decode(&input)

	token, err := k.LoginUser(input.Username, input.Password)
	if err != nil {
		log.Println("Error logging in with user:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	defer c.Request.Body.Close()

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, PassItResponseBody{
		Code: codes.UserCreatedSuccessfully,
		Data: nil,
	},
	)

}

func (s *Server) CreateUserHandler(c *gin.Context) {
	var input CreateUserRequestBody
	var keycloackClient keyclock.KeycloakClient
	k := keycloackClient.NewKeycloakClient()

	json.NewDecoder(c.Request.Body).Decode(&input)

	user := input.User

	keyclockUserID, err := k.CreateUser(&user, input.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in Keycloak"})
		return
	}

	user.KeycloackID = keyclockUserID
	err = s.db.CreateUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, PassItResponseBody{
		Code: codes.UserCreatedSuccessfully,
		Data: user,
	},
	)
}

func (s *Server) FindUserByIdHandler(c *gin.Context) {
	var input FindUserByIdRequestBody

	json.NewDecoder(c.Request.Body).Decode(&input)
	user, err := s.db.FindUserById(input.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, user)
}

func (s *Server) FindUserByEmailHandler(c *gin.Context) {
	var input FindUserByEmailRequestBody

	json.NewDecoder(c.Request.Body).Decode(&input)
	user, err := s.db.FindUserByEmail(input.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, user)
}

func (s *Server) UpdateUserByIdHandler(c *gin.Context) {
	var user models.User

	json.NewDecoder(c.Request.Body).Decode(&user)
	err := s.db.UpdateUserById(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, user)
}

func (s *Server) GetAllUsersHandler(c *gin.Context) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, users)
}
