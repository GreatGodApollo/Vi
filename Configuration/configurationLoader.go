package Configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Bot struct {
		Token    string   `json:"token"`
		Prefixes []string `json:"prefixes"`
		Owners   []string `json:"owners"`
	} `json:"bot"`
}

func LoadConfiguration(file string) Configuration {
	var config Configuration
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)
	return config
}
