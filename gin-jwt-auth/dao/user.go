package dao

import (
	"errors"
	"fmt"

	"../common"
	"../db"
	"../model"
	"../utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]model.User, error) {
	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	var users []model.User
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (model.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return model.User{}, err
	}

	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	var user model.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (u *User) Login(email string, password string) (model.User, error) {
	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	var user model.User
	err := collection.Find(bson.M{"email": email}).One(&user)
	fmt.Println(password)
	fmt.Println(user.Password)
	fmt.Println(u.utils.ValidatePassword(password, user.Password))
	if !u.utils.ValidatePassword(password, user.Password) {
		err = errors.New(common.ErrInvalidCredentials)
	}
	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user model.User) error {
	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	// Encrypt password
	user.Password = u.utils.EncryptPassword(user.Password)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
func (u *User) Delete(user model.User) error {
	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (u *User) Update(user model.User) error {
	sessionCopy := db.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Database.Databasename).C(common.ColUsers)

	err := collection.UpdateId(user.ID, &user)
	return err
}
