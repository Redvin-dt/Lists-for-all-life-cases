package entities

type List struct {
	Values []string
}

type ListMongo struct {
	Values []string `bson:"Values"`
}
