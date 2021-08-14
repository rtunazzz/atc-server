package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadConfig(filename string) (config map[string]SiteConfig, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	return
}
