package Config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

var envMap map[string]string
var err error

func InitEnv() {
	if envMap == nil {
		envMap, err = godotenv.Read()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func GetEnvKey(key string) string {
	return envMap[key]
}

func GetEnvKeyInt(key string) int {
	value, err := strconv.Atoi(envMap[key])
	if err != nil {
		log.Fatal("Error parsing value to int")
	}
	return value
}

func GetEnvKeyBool(key string) bool {
	value, err := strconv.ParseBool(envMap[key])
	if err != nil {
		log.Fatal("Error parsing value to int")
	}
	return value
}
