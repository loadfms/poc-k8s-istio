package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	const domain = "identity"
	// Create an Echo instance
	e := echo.New()

	// Define your JWT secret key
	secretKey := []byte("your-secret-key-here")

	// Generate the JWKS
	jwks, err := generateJWKS(secretKey)
	if err != nil {
		fmt.Printf("Error generating JWKS: %v\n", err)
		return
	}

	// Middleware for logging and recovery
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group(domain)

	// Route handler for the root path "/"
	apiGroup.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("olar %s", c.Request().URL)
		return c.JSON(http.StatusOK, map[string]interface{}{"message": message})
	})

	// Route handler to serve the JWKS
	apiGroup.GET("/.well-known/jwks.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, jwks)
	})

	// Route handler for "/signin"
	apiGroup.GET("/signin", func(c echo.Context) error {
		// Generate a JWT token
		jwt := generateToken(secretKey)
		return c.JSON(http.StatusOK, jwt)
	})

	// Start the HTTP server
	port := ":8082" // Replace with your desired port
	fmt.Printf("Server listening on port %s...\n", port)
	e.Start(port)
}

func generateJWKS(secretKey []byte) (*customJWKS, error) {
	// Create a new JSON Web Key (JWK) using your secret key
	key, err := jwk.New(secretKey)
	if err != nil {
		return nil, err
	}

	// Create a JSON Web Key Set (JWKS) containing the JWK
	jwks := &customJWKS{
		Keys: []jwk.Key{key},
	}

	return jwks, nil
}

func generateToken(secretKey []byte) map[string]interface{} {
	// Define the JWT claims (payload)
	claims := jwt.MapClaims{
		"sub":  "1234567890",                            // Subject
		"name": "John Doe",                              // Custom claims
		"iat":  time.Now().Unix(),                       // Issued At (Unix timestamp)
		"exp":  time.Now().Add(time.Minute * 15).Unix(), // Expiration time (Unix timestamp)
		"iss":  "https://loadfms.local",
	}

	// Define the JWT signing method and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return nil
	}

	// Create a map to represent the JWT response
	jwtResponse := map[string]interface{}{
		"token": tokenString,
	}

	return jwtResponse
}

// Define a custom JWKS structure
type customJWKS struct {
	Keys []jwk.Key `json:"keys"`
}
