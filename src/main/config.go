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
