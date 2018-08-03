package main

import (
	"os"
	"path/filepath"

	"../config"
	"../server"
)

func main() {
	d, _ := os.Getwd()
	d = filepath.FromSlash(d)
	configuration := config.LoadConfig(filepath.Join(d, "config.json"))
	server.Start(configuration)
}
