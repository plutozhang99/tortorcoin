package auth

import (
	"errors"
	"time"
	"tortorCoin/pkg/config"

	"github.com/dgrijalva/jwt-go"
)

// Claims struct to add claims to the token
type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

// Initialize the jwtKey variable
var jwtKey []byte
var experiationTime time.Duration

// InitConfig initializes the configuration for the JWT by loading it from the config package.
func InitConfig() error {
	// Load the configuration using LoadConfig function from the config package
	cfg, err := config.LoadConfig() // Use the actual path to your config directory
	if err != nil {
		return err
	}
	jwtKey = []byte(cfg.JWT.SecretKey) // Set the jwtKey to the secret key from the config
	experiationTime = time.Duration(cfg.JWT.ExpirationTime)
	return nil
}

// GenerateJWT generates a new JWT token for a given Account
func GenerateJWT(account string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(experiationTime)
	// Create the JWT claims including the Account and expiry time
	claims := &Claims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates the JWT token and returns the Account if it's valid
func ValidateJWT(tokenString string) (string, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return "", err
		}
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Return the Account from the token's claims
	return claims.Account, nil
}
