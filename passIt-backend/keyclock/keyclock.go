package keyclock

import (
	"context"
	"crypto/tls"
	"log"
	"passIt/internal/models"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-resty/resty/v2"
)

var KeycloackHost = "https://localhost:8443"
var KeycloackRealm = "passit"
var KeycloackAdminUsername = "admin"
var KeycloackAdminPassword = "admin"

type KeycloakClient struct {
	client      *gocloak.GoCloak
	restyClient *resty.Client
	ctx         context.Context
}

func (k *KeycloakClient) NewKeycloakClient() *KeycloakClient {
	restyClient := resty.New()
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	client := gocloak.NewClient(KeycloackHost)
	client.SetRestyClient(restyClient)

	return &KeycloakClient{
		client:      client,
		restyClient: restyClient,
		ctx:         context.Background(),
	}
}

func (k *KeycloakClient) GetUserInfo(token string) (*gocloak.UserInfo, error) {
	userInfo, err := k.client.GetUserInfo(k.ctx, token, KeycloackRealm)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (k *KeycloakClient) CreateUser(user *models.User, password string) (string, error) {
	log.Println(&k)
	token, err := k.client.LoginAdmin(k.ctx, KeycloackAdminUsername, KeycloackAdminPassword, KeycloackRealm)
	if err != nil {
		log.Println("Error logging in to Keycloak:", err)
		return "", err
	}
	gocloakUser := gocloak.User{
		Username:  gocloak.StringP(user.Username),
		FirstName: gocloak.StringP(user.FirstName),
		LastName:  gocloak.StringP(user.LastName),
		Email:     gocloak.StringP(user.Email),
		Enabled:   gocloak.BoolP(true), // Ensure the user is enabled
	}
	UserID, err := k.client.CreateUser(k.ctx, token.AccessToken, KeycloackRealm, gocloakUser)
	if err != nil {
		log.Println("Error creating user in Keycloak:", err)
		return "", err
	}

	// Set the user's password
	err = k.client.SetPassword(k.ctx, token.AccessToken, UserID, KeycloackRealm, password, false)
	if err != nil {
		log.Println("Error setting password for user in Keycloak:", err)
		return "", err
	}

	// Handle custom fields separately if needed
	// Example: Save DOB, PhoneNumber, Address, IsActive, IsAdmin in a custom database or another service
	return UserID, nil
}
