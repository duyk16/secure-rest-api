package model

import (
	"context"
	"time"

	"github.com/duyk16/secure-rest-api/storage"
	"github.com/duyk16/secure-rest-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func (u *User) InsertUser() (err error) {
	// Create new ObjectID
	u.ID = primitive.NewObjectID()

	// Get current time
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// Hash password
	u.Password = util.HashAndSaltPassword(u.Password)

	_, err = storage.User.InsertOne(context.Background(), u)

	return err
}

func (u *User) GetUserById() (err error) {
	result := storage.User.FindOne(
		context.Background(),
		bson.D{{"_id", u.ID}},
	)
	return result.Decode(u)
}

func (u *User) GetUserByEmail() (err error) {
	result := storage.User.FindOne(
		context.Background(),
		bson.D{{"email", u.Email}},
	)

	return result.Decode(u)
}
