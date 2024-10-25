package entities

type User struct {
	Login          string `json:"login"`
	HashedPassword string `json:"password"`
}
