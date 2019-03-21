package models

type gender string

const (
	male   gender = "Male"
	female gender = "Female"
)

// User struct
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Age       int8   `json:"age,omitempty"`
	Gender    gender `json:"gender,omitempty"`
}
