package main

import (
	"secure-rest-api/config"
	"secure-rest-api/router"
	"secure-rest-api/storage"
)

func main() {
	config.Init()
	storage.Init()
	router.Init()
}
