package controller

import (
	"encoding/json"
	"net/http"

	"secure-rest-api/model"
	u "secure-rest-api/util"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var postUser model.User
	err := json.NewDecoder(r.Body).Decode(&postUser)

	if err != nil {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Body error",
		})
		return
	}

	if postUser.Email == "" || postUser.Password == "" {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Email or password is not valid",
		})
		return
	}

	user := model.User{
		Email:    postUser.Email,
		Password: postUser.Password,
	}

	err = user.GetUserByEmail()

	if err != nil {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Not found email",
		})
		return
	}

	ok := u.ComparePasswords(user.Password, postUser.Password)

	if !ok {
		u.JSON(w, 400, u.T{
			"status":  "error",
			"message": "Password is not correct",
		})
		return
	}

	token := u.GenerateToken(user.ID, user.Email)
	u.JSON(w, 200, u.T{
		"status": "success",
		"data": u.T{
			"id":    user.ID,
			"email": user.Email,
			"token": token,
		},
	})
	return
}
