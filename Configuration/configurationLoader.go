/*
 * Vi - A Discord Bot written in Go
 * Copyright (C) 2019  Brett Bender
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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
	Database struct {
		Connection string `json:"connection"`
	} `json:"db"`
	Miscellaneous struct {
		ColorEnabled      bool   `json:"colorEnabled"`
		SuggestionChannel string `json:"suggestions"`
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
