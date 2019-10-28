package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"secure-rest-api/model"
	"secure-rest-api/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var postUser model.User
	err := json.NewDecoder(r.Body).Decode(&postUser)

	if err != nil {
		util.SetResponseHeader(w, 400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Body is not valid",
		})
		return
	}

	if postUser.Email == "" || postUser.Password == "" {
		util.SetResponseHeader(w, 400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Email or Password is not valid",
		})
		return
	}

	hash := util.HashAndSaltPassword(postUser.Password)

	log.Println("HASH", hash)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": postUser,
	})
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	result := util.ComparePasswords("$2a$04$c.VnJyCq3kb7D.Zz.FVSqurythvTCKmBB6zcQQtWqvRDsD6GVgFym", "123456")
	log.Println(result)

	w.Write([]byte("oke"))
}
