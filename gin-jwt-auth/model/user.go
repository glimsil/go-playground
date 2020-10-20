package model

import (
	"errors"

	"../common"
	"gopkg.in/mgo.v2/bson"
)

// User information
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Email    string        `bson:"email" json:"email" example:"glimsil@email.com"`
	Name     string        `bson:"name" json:"name" example:"gustavo"`
	Password string        `bson:"password" json:"password" example:"randompass123"`
	Roles    []string      `bson:"roles" json:"roles" example:"['USER']"`
}

// AddUser information
type AddUser struct {
	Email    string `json:"email" example:"glimsil@email.com"`
	Name     string `json:"name" example:"Gustavo"`
	Password string `json:"password" example:"randompass123"`
}

// Validate user
func (a AddUser) Validate() error {
	switch {
	case len(a.Email) == 0:
		return errors.New(common.ErrEmailEmpty)
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}
