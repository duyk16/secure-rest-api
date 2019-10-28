package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Threads int     `json:"threads"`
	Name    string  `json:"name"`
	Storage Storage `json:"storage"`
	JWTKey  string  `json:"jwt_key"`
}

type Storage struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

var ServerConfig Config

func Init() {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&ServerConfig); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}
