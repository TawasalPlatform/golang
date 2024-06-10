package tawasal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// GetUser extracts and decodes the user information from a provided cookie.
// Parameters:
// - cookie: A raw string representing the cookie from which user information is to be extracted.
// Returns: An object containing the user information.
func GetUser(cookie string) (*User, error) {
	var tawasal User

	if cookie == "" {
		return &tawasal, nil
	}

	decoded, err := DecodeBase64(cookie)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(decoded), &tawasal)
	if err != nil {
		return nil, err
	}

	return &tawasal, nil
}

// GetAuthorization generates an authorization token from the provided cookie.
// Parameters:
// - cookie: A raw string representing the cookie from which the authorization token is to be extracted.
// Returns: A base64 encoded string representing the authorization token, or an error if the token is not available.
func GetAuthorization(cookie string) (string, error) {
	tawasal, err := GetUser(cookie)
	if err != nil {
		return "", err
	}

	if tawasal.UserToken == "" {
		return "", nil
	}

	tokenParts := strings.Split(tawasal.UserToken, ":")
	if len(tokenParts) != 4 {
		return "", fmt.Errorf("invalid token format")
	}

	authToken := fmt.Sprintf("%d:%s:%s:%s:%s", tawasal.UserID, tokenParts[0], tokenParts[1], tokenParts[2], tokenParts[3])
	return base64.StdEncoding.EncodeToString([]byte(authToken)), nil
}

// GetDeviceToken extracts the device token from the provided cookie.
// Parameters:
// - cookie: A raw string representing the cookie from which the device token is to be extracted.
// Returns: A string representing the device token, or an error if the token is not available.
func GetDeviceToken(cookie string) (string, error) {
	tawasal, err := GetUser(cookie)
	if err != nil {
		return "", err
	}

	if tawasal.UserToken == "" {
		return "", nil
	}

	tokenParts := strings.Split(tawasal.UserToken, ":")
	if len(tokenParts) < 3 {
		return "", fmt.Errorf("invalid token format")
	}

	return tokenParts[2], nil
}
