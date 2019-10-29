package main

import (
	"github.com/duyk16/secure-rest-api/config"
	"github.com/duyk16/secure-rest-api/router"
	"github.com/duyk16/secure-rest-api/storage"
)

func main() {
	config.Init()
	storage.Init()
	router.Init()
}
