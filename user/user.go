package user

//easyjson:json
type User struct {
	Name     string   `json:"name"`
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
}
