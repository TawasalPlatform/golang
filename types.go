package tawasal

// Language type represents supported languages.
type Language string

const (
	English    Language = "en"
	Arabic     Language = "ar"
	Spanish    Language = "es"
	Persian    Language = "fa"
	French     Language = "fr"
	Indonesian Language = "id"
	Russian    Language = "ru"
	Turkish    Language = "tr"
	Hindi      Language = "hi"
	Urdu       Language = "ur"
)

// Platform type represents supported platforms.
type Platform string

const (
	IOS     Platform = "ios"
	Android Platform = "android"
)

// User struct represents the user information.
type User struct {
	UserID       int      `json:"userId"`
	UserToken    string   `json:"userToken,omitempty"`
	FirstName    string   `json:"firstName,omitempty"`
	LastName     string   `json:"lastName,omitempty"`
	UserNickname string   `json:"userNickname,omitempty"`
	Language     Language `json:"language"`
	Platform     Platform `json:"platform"`
	Version      string   `json:"version"`
}
