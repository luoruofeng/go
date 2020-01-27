package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	LbAddr string `json:"lb_addr"`
}

type Test struct {
	Name string `json:"lb_addr"`
}

var configuration *Configuration
var test *Test

func init() {
	file, _ := os.Open("../conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)

	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {

		log.Fatal("../conf.json read failed!!!!!!")
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LbAddr
}
