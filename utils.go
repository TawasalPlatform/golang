package tawasal

import "encoding/base64"

// DecodeBase64 decodes a base64 encoded string similar to the JavaScript implementation.
func DecodeBase64(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
