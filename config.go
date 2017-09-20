package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string
}

var GlobalConfig Config

func LoadConfig() {
	data,err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	GlobalConfig = Config{}
	err = json.Unmarshal(data,&GlobalConfig)
	if err != nil {
		log.Fatal(err)
	}
}