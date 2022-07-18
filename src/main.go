package main

import (
	"os"

	"github.com/daniel5268/go-meye/src/api"
	"github.com/daniel5268/go-meye/src/config"
)

func main() {
	env := os.Getenv("GO_ENV")
	if env != "production" {
		config.LoadConfig(config.Dev)
	} else {
		config.LoadConfig(config.Prod)
	}

	api.NewApp().StartApp()
}
