package router

import (
	"log"
	"net/http"
	"secure-rest-api/controller"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/api/auth/login", controller.PostLogin).Methods("POST")
	r.HandleFunc("/api/user", controller.CreateUser).Methods("POST")

	r.HandleFunc("/api/user/{userId}", controller.GetUserById).Methods("GET")

	listen(r, "8000")
}

func listen(r *mux.Router, port string) {
	log.Printf("Server listening on port %v...\n", port)

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Println("Serve server fail", err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {

}
