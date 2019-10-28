package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"secure-rest-api/model"
	u "secure-rest-api/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var postUser model.User
	err := json.NewDecoder(r.Body).Decode(&postUser)

	if err != nil {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Body is not valid",
		})
		return
	}

	if postUser.Email == "" || postUser.Password == "" {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Email or Password is not valid",
		})
		return
	}

	err = postUser.InsertUser()

	if err != nil {
		mongoErr, ok := err.(mongo.WriteException)

		if ok && mongoErr.WriteErrors[0].Code == 11000 {
			u.JSON(w, 400, u.T{
				"status":  "error",
				"code":    11000,
				"message": "Email was used before",
			})
			return
		}

		u.JSON(w, 500, u.T{
			"status":  "error",
			"message": "Insert user fail",
		})
		return
	}

	u.JSON(w, 201, u.T{
		"status":  "success",
		"message": postUser,
	})
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	result := u.ComparePasswords("$2a$04$c.VnJyCq3kb7D.Zz.FVSqurythvTCKmBB6zcQQtWqvRDsD6GVgFym", "123456")
	log.Println(result)

	w.Write([]byte("oke"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["userId"]
	userID, err := primitive.ObjectIDFromHex(idString)

	if err != nil {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "User ID is not valid",
		})
	}

	var user model.User
	user.ID = userID
	err = user.GetUserById()

	if err != nil {
		log.Printf("%T %v", err, err)
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Not found User ID",
		})
		return
	}

	u.JSON(w, 200, u.T{
		"status": "success",
		"data": u.T{
			"id":    user.ID,
			"email": user.Email,
		},
	})
	return
}
