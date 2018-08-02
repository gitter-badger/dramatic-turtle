package main

import (
	"encoding/json"
	"io/ioutil"

	"../core"
)

// Config struct
type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Storage struct {
		MongoDB struct {
			Active       bool `json:"active"`
			ReadEnabled  bool `json:"read_enabled"`
			WriteEnabled bool `json:"write_enabled"`
			Connection   struct {
				URL   string `json:"url"`
				Name  string `json:"name"`
				Tasks string `json:"tasks"`
				Logs  string `json:"logs"`
			} `json:"connection"`
		} `json:"mongodb"`
	} `json:"storage"`
}

// LoadConfig func
func LoadConfig(fileName string) Config {
	var data Config
	file, err := ioutil.ReadFile(fileName)
	core.CheckErr(err)
	err = json.Unmarshal(file, &data)
	core.CheckErr(err)

	return data
}
