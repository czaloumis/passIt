package auth

import (
	"context"
	"passIt/internal/utils"

	// "crypto/rsa"

	// "crypto/tls"
	// "crypto/x509"

	// // "encoding/base64"
	// "encoding/json"
	// "encoding/pem"
	// "errors"
	"fmt"
	// "log"

	// "net/http"
	// "passIt/internal/models"
	// "sync"

	// "github.com/Nerzal/gocloak/v13"
	// "github.com/go-resty/resty/v2"
	// "github.com/golang-jwt/jwt/v5"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Config struct {
	BaseURL      string // Authorization base url
	ClientID     string // client id oauth
	RedirectURL  string // valid redirect url
	ClientSecret string // keycloak client secret
	Realm        string // keycloak realm
}

// Client struct holds all components needed for authentication
type Client struct {
	Provider *oidc.Provider        // Handles OIDC protocol operations with Keycloak
	OIDC     *oidc.IDTokenVerifier // Verifies JWT tokens from Keycloak
	Oauth    oauth2.Config         // Manages OAuth2 flow (authorization codes, tokens)
}

func New(ctx context.Context, config *Config) (*Client, error) {
	// Create an insecure HTTP client (for self-signed certs only)
	// insecureHttpClient := &http.Client{
	// 	Timeout: 10 * time.Second,
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 👈 disables cert check
	// 	},
	// }

	// Use the custom HTTP client in the context
	insecureCtx := utils.InsecureHttpContext(ctx)
	// Construct the provider URL using Keycloak realm
	providerURL := fmt.Sprintf("%s/realms/%s", config.BaseURL, config.Realm)

	provider, err := oidc.NewProvider(insecureCtx, providerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}

	// Create ID token verifier
	verifier := provider.Verifier(&oidc.Config{
		ClientID: config.ClientID,
	})

	// Configure an OpenID Connect aware OAuth2 client with specific scopes:
	// - oidc.ScopeOpenID: Required for OpenID Connect authentication, provides subject ID (sub)
	// - "roles": Keycloak-specific scope to get user roles in the token
	oauth2Config := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes: []string{
			oidc.ScopeOpenID, // Required for OIDC authentication
			"roles",          // Request user roles from Keycloak
		},
	}

	// Return initialized client with all required components
	return &Client{
		// oauth2Config: Used for OAuth2 operations like:
		// - Generating login URL (AuthCodeURL)
		// - Exchanging auth code for tokens (Exchange)
		// - Managing token refresh
		Oauth: oauth2Config,

		// verifier: Used to validate tokens:
		// - Verifies JWT signature
		// - Validates token claims (exp, iss, aud)
		// - Extracts user information
		OIDC: verifier,

		// provider: Keycloak OIDC provider that:
		// - Provides endpoint URLs (auth, token)
		// - Handles OIDC protocol details
		// - Manages provider metadata
		Provider: provider,
	}, nil
}

// AuthCodeURL generates the login URL for OAuth2 authorization code flow.
// It returns a URL that the user should be redirected to for authentication.
// The state parameter is a random string that will be validated in the callback
// to prevent CSRF attacks.
func (c *Client) AuthCodeURL(state string) string {
	return c.Oauth.AuthCodeURL(state)
}
