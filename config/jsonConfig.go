package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database string
	Port string
}

func GetConfig()  Config{
	config := Config{
		Database: "blogs.db",
		Port: ":8080",
	}

	configFile, err := os.Open("../cmd"+"./config/jsonConfig.go")
	if err != nil {
		log.Fatalln("",err)
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}