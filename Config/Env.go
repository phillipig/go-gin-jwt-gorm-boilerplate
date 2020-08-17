package config

import (
	"log"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type EnvMap map[string]string

var once sync.Once
var instance EnvMap
var err error

func NewEnv() EnvMap {
	once.Do(func() {
		instance, err = godotenv.Read()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	})
	return instance
}

func (env EnvMap) GetEnvKey(key string) string {
	return env[key]
}

func (env EnvMap) GetEnvKeyInt(key string) int {
	value, err := strconv.Atoi(env[key])
	if err != nil {
		log.Fatal("Error parsing value to int")
	}
	return value
}

func (env EnvMap) GetEnvKeyBool(key string) bool {
	value, err := strconv.ParseBool(env[key])
	if err != nil {
		log.Fatal("Error parsing value to int")
	}
	return value
}
