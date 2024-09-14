package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database string
	Port string
	Secret string
}

func GetConfig()  Config{
	var config Config

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("",err)
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}