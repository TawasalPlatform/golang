<div align="center">
  <a href="https://tawasal.ae/">
    <img src="https://tawasal.ae/tawasal_logo_full.png" width="300" height="auto" alt="Tawasal"/>
  </a>
</div>
<hr />

<div align="center">
    <p align="center">
        <a href="https://platform.tawasal.ae"><b>check our Documentation 👉 platform.tawasal.ae</b></a><br />
    </p>
</div>
<hr />

# Tawasal SDK for Go

The Tawasal SDK for Go provides a set of utilities to interact with the Tawasal platform. This SDK allows you to extract and decode user information from a cookie, generate authorization tokens, and extract device tokens.

## Installation

To install the Tawasal SDK, use the following command:

```sh
go get github.com/TawasalPlatform/golang
```

## Usage

### Import the Package

In your Go code, import the Tawasal SDK package:

```go
import tawasal "github.com/TawasalPlatform/golang"
```

### Functions Provided by the SDK

#### `GetUser`

Extracts and decodes the user information from a provided cookie.

```go
import (
    tawasal "github.com/TawasalPlatform/golang"
)

func main() {
    cookie := "your_encoded_cookie_here"
    user, err := tawasal.GetUser(cookie)
    if err != nil {
        log.Fatalf("Error getting user: %v", err)
    }
    fmt.Printf("User: %+v\n", user)
}
```

#### `GetAuthorization`

Generates an authorization token from the provided cookie.

```go
import (
    tawasal "github.com/TawasalPlatform/golang"
)

func main() {
    cookie := "your_encoded_cookie_here"
    authToken, err := tawasal.GetAuthorization(cookie)
    if err != nil {
        log.Fatalf("Error getting authorization token: %v", err)
    }
    fmt.Printf("Authorization Token: %s\n", authToken)
}
```

#### `GetDeviceToken`

Extracts the device token from the provided cookie.

```go
package main

import (
	"fmt"
	tawasal "github.com/TawasalPlatform/golang"
	"log"
)

func main() {
	cookie := "your_encoded_cookie_here"
	deviceToken, err := tawasal.GetDeviceToken(cookie)
	if err != nil {
		log.Fatalf("Error getting device token: %v", err)
	}
	fmt.Printf("Device Token: %s\n", deviceToken)
}
```

#### CheckSignature( userId: number, authKeyId: string, deviceToken: string, signatureBase64: string, publicKey: string)
This function verifies user.

- **Parameters**:
    - `userId`: id of the tawasal user,
    - `authKeyId`: key of authorisation, second part of user token,
    - `deviceToken`: the token describing session on given device,
    - `signatureBase64`: first part od user token,
    - `publicKey`: the key that will be obtained in Dev Management
- **Returns**: A boolean that says if session are legit.

**Example**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	tawasal "github.com/TawasalPlatform/golang"
	"strings"
)

// User struct to hold user data
type User struct {
	UserID       int    `json:"userId"`
	UserToken    string `json:"userToken"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	UserNickname string `json:"userNickname"`
	Language     string `json:"language"`
	Platform     string `json:"platform"`
	Version      string `json:"version"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the cookie
		cookie, err := r.Cookie("tawasal")
		if err != nil {
			http.Error(w, "Cookie not found", http.StatusBadRequest)
			return
		}

		// Get the user information from the cookie
		user, err := tawasal.GetUser(cookie.Value)
		if err != nil {
			http.Error(w, "Failed to get user from cookie", http.StatusInternalServerError)
			return
		}

		if user.UserToken != "" {
			tokenParts := strings.Split(user.UserToken, ":")
			if len(tokenParts) != 4 {
				http.Error(w, "Invalid token format", http.StatusBadRequest)
				return
			}

			signature := tokenParts[0]
			authKeyId := tokenParts[1]
			deviceToken := tokenParts[2]
			publicKey := `-----BEGIN PUBLIC KEY----------END PUBLIC KEY-----` // will be obtained at Dev Managment

			result, err := tawasal.CheckSignature(user.UserID, authKeyId, deviceToken, signature, publicKey)
			if err != nil {
				http.Error(w, "Failed to check signature", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Signature valid: %v", result)
		} else {
			http.Error(w, "User token not found", http.StatusBadRequest)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
```

## Example

Here's a complete example demonstrating how to use the Tawasal SDK in a Go application:

```go
package main

import (
	"fmt"
	"log"
	tawasal "github.com/TawasalPlatform/golang"
)

func main() {
	cookie := "your_encoded_cookie_here"

	// Get user information
	user, err := tawasal.GetUser(cookie)
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	fmt.Printf("User: %+v\n", user)

	// Get authorization token
	authToken, err := tawasal.GetAuthorization(cookie)
	if err != nil {
		log.Fatalf("Error getting authorization token: %v", err)
	}
	fmt.Printf("Authorization Token: %s\n", authToken)

	// Get device token
	deviceToken, err := tawasal.GetDeviceToken(cookie)
	if err != nil {
		log.Fatalf("Error getting device token: %v", err)
	}
	fmt.Printf("Device Token: %s\n", deviceToken)
}
```

## API Reference

### `GetUser`

Extracts and decodes the user information from a provided cookie.

#### Parameters

- `cookie` : A raw string representing the cookie from which user information is to be extracted.

#### Returns

- An object containing the user information.
- An error, if any.

### `GetAuthorization`

Generates an authorization token from the provided cookie.

#### Parameters

- `cookie` : A raw string representing the cookie from which the authorization token is to be extracted.

#### Returns

- A base64 encoded string representing the authorization token, or an error if the token is not available.

### `GetDeviceToken`

Extracts the device token from the provided cookie.

#### Parameters

- `cookie` : A raw string representing the cookie from which the device token is to be extracted.

#### Returns

- A string representing the device token, or an error if the token is not available.

## License

This project is licensed under the MIT License.
