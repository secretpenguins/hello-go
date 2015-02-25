package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)
type Config struct {
	Development ConfigSection
	Production ConfigSection
}

type ConfigSection struct {
	DbUser string
	DbPassword string
	DbHost string
	DbPort string
	DbDatabase string
	MemcachePath string
}

var globalConfig Config

func init() {
	Setup()
}

func Setup() {
	//initializeConfig()
	log.Println("Init Config")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error", err)
	}
	//decoder := json.NewDecoder(file)
	//err2 := decoder.Decode(&config)
	/*if err2 != nil {
  		fmt.Println("error:", err)
	}*/

	fmt.Println("JSON: %s", string(file))

	json.Unmarshal(file, &globalConfig)
	fmt.Println(globalConfig.Development)
}

func GetConfig() ConfigSection {
	if os.Getenv("GO_ENV") == "production" {
		return globalConfig.Production
	} else {
		return globalConfig.Development
	}
}

