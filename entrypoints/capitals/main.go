package main

import (
	"go-api-template/cmd"
	"go-api-template/config"
)

func main() {

	config.LoadEnv(config.RootPath() + "/config/.env")

	cmd.Execute()
}
