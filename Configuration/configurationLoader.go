package Configuration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Configuration struct {
	Bot struct {
		Token          string   `json:"token"`
		Prefixes       []string `json:"prefixes"`
		Owners         []string `json:"owners"`
		Statuses       []string `json:"statuses"`
		StatusInterval string   `json:"statusInterval"`
	} `json:"bot"`
	Miscellaneous struct {
		ColorEnabled bool `json:"colorEnabled"`
	} `json:"miscellaneous"`
}

func LoadConfiguration(file string, log *logrus.Logger) Configuration {
	var config Configuration
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)
	return config
}
