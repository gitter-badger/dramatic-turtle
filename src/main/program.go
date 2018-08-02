package main

import (
	"os"
	"path/filepath"

	"../server"
)

func main() {
	d, _ := os.Getwd()
	d = filepath.FromSlash(d)
	config := LoadConfig(filepath.Join(d, "config.json"))
	server.Start(config.Server.Port)
}
