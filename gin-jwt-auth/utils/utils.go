package utils

import (
	"errors"
	"time"

	"../common"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	UserId string   `json:"userId"`
	Email  string   `json:"email"`
	Roles  []string `json:"role"`
	jwt_lib.StandardClaims
}

type Utils struct {
}

// GenerateJWT generates token from the given information
func (u *Utils) GenerateJWT(userId string, email string, role []string) (string, error) {
	claims := SdtClaims{
		userId,
		email,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    common.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))

	return tokenString, err
}

// ValidateObjectID checks the given ID if it's an object id or not
func (u *Utils) ValidateObjectID(id string) error {
	if bson.IsObjectIdHex(id) != true {
		return errors.New(common.ErrNotObjectIDHex)
	}

	return nil
}

// Encrypt a string using bcrypt
func (u *Utils) EncryptPassword(str string) string {
	bytePassword := []byte(str)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// Encrypt a string using bcrypt
func (u *Utils) ValidatePassword(original string, hashed string) bool {
	byteOriginal := []byte(original)
	byteHashed := []byte(hashed)

	// Hashing the password with the default cost of 10
	err := bcrypt.CompareHashAndPassword(byteHashed, byteOriginal)
	return err == nil
}
