package main

import (
	"github.com/vanya/backend/internal/app"
)

const configPath = "configs"

func main() {
	app.Run(configPath)
}
