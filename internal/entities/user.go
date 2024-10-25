package entities

type User struct {
	Login          string `json:"login"`
	HashedPassword string `json:"password"`
}

type MongoUser struct {
	Login          string `bson:"login"`
	HashedPassword string `bson:"password"`
}
