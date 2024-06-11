package tawasal

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
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

// CheckSignature verifies the user session based on provided parameters.
// Parameters:
// - userId: ID of the tawasal user
// - authKeyId: Key of authorization, second part of user token
// - deviceToken: The token describing session on given device
// - signatureBase64: First part of user token
// - publicKey: The key obtained in Dev Management
// Returns: A boolean indicating if the session is legitimate and an error if any.
func CheckSignature(userId int, authKeyId, deviceToken, signatureBase64, publicKey string) (bool, error) {
	// Decode the base64 signature
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	// Decode the public key
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return false, errors.New("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("public key is not of type RSA")
	}

	// Create the data to be verified
	data := []byte(fmt.Sprintf("%d%s%s", userId, authKeyId, deviceToken))

	// Create a SHA256 hash of the data
	hash := sha256.New()
	hash.Write(data)
	hashed := hash.Sum(nil)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, hashed, signature)
	if err != nil {
		return false, nil
	}

	return true, nil
}
