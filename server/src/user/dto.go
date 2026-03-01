package user

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"min=3,max=50"`
	Username string `json:"username" validate:"min=3,max=20"`
}

type GetUserResponse struct {
	Id       bson.ObjectID `json:"id"`
	Name     string        `json:"name"`
	Username string        `json:"username"`
}

func (g *GetUserResponse) Set(user *User) {
	g.Id = user.Id
	g.Name = user.Name
	g.Username = user.Username
}
