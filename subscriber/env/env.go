package env

import (
	"encoding/json"
	"log"
	"os"
)

type Env struct {
	Redis EnvRedis `json:"redis"`
	SMTP  EnvSMTP  `json:"smtp"`
}

type EnvRedis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type EnvSMTP struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

var ENV Env

func NewEnvironment() {
	env, err := os.ReadFile("./env.json")
	if err != nil {
		log.Fatalln("could not load env")
	}
	err = json.Unmarshal(env, &ENV)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
