package entities

type User struct {
	Username string   `json:"username"`
	Password string   `json:"-"`
	Roles    []string `json:"roles"`
}
