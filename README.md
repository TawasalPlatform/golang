#
<div align="center">
  <a href="https://tawasal.ae/">
    <img src="https://tawasal.ae/tawasal_logo_full.png" width="300" height="auto" alt="Tawasal"/>
  </a>
</div>

<hr />

<p align="center">
<a href="https://platform.tawasal.ae"><b>check our Documentation 👉 platform.tawasal.ae</b></a><br />
</p>

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
import "github.com/TawasalPlatform/golang"
```

### Functions Provided by the SDK

#### `GetUser`

Extracts and decodes the user information from a provided cookie.

```go
import (
	"github.com/TawasalPlatform/golang"
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
	"github.com/TawasalPlatform/golang"
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
import (
	"github.com/TawasalPlatform/golang"
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

## Example

Here's a complete example demonstrating how to use the Tawasal SDK in a Go application:

```go
package main

import (
	"fmt"
	"log"
	"github.com/TawasalPlatform/golang"
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

- `cookie`: A raw string representing the cookie from which user information is to be extracted.

#### Returns

- An object containing the user information.
- An error, if any.

### `GetAuthorization`

Generates an authorization token from the provided cookie.

#### Parameters

- `cookie`: A raw string representing the cookie from which the authorization token is to be extracted.

#### Returns

- A base64 encoded string representing the authorization token, or an error if the token is not available.

### `GetDeviceToken`

Extracts the device token from the provided cookie.

#### Parameters

- `cookie`: A raw string representing the cookie from which the device token is to be extracted.

#### Returns

- A string representing the device token, or an error if the token is not available.

## License

This project is licensed under the MIT License.

---

This README provides an overview of how to use the Tawasal SDK in a Go application. Adjust the import paths and package names as necessary to match your project's structure.
