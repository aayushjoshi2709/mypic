package user

type User struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	UserName string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
