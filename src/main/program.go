package main

import (
	"os"
	"path/filepath"

	"../config"
	"../server"
)

func main() {
	print(os.Args)
	d := filepath.Dir(os.Args[0])
	d = filepath.FromSlash(d)
	configuration := config.LoadConfig(filepath.Join(d, "config.json"))
	server.Start(configuration)
}
