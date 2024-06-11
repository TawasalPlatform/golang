package tawasal

import (
	"encoding/base64"
	"strings"
)

// DecodeBase64 decodes a base64 encoded string similar to the JavaScript implementation.
func DecodeBase64(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// FixBase64Padding ensures the base64 string has the correct padding
func fixBase64Padding(encoded string) string {
	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}
	return encoded
}
