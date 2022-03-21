package structs

type User struct {
	Username string `json:"username" bson:"username"`
	APIKey   string `json:"api-key" bson:"api-key"`
	IsAdmin  bool   `json:"is-admin" bson:"is-admin"`
}
