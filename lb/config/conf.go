package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	LbAddr string `json:"lb_addr"`
}

var configuration *Configuration

//ConfigurationPath is json configuration file path.It should be changed when producation env.
const ConfigurationPath string = "/Users/luoruofeng/Desktop/go/lb/config/conf.go"

func init() {
	file, _ := os.Open("/Users/luoruofeng/Desktop/go/lb/conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		log.Fatal("../conf.json read failed!!  ", err)
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LbAddr
}
