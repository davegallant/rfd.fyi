package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	utils "github.com/davegallant/rfd-fyi/pkg/utils"
)

// @title           RFD FYI API
// @version         1.0
// @description     An API for an issue tracking service
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	a := &App{}

	httpPort := utils.GetEnv("HTTP_PORT", "8080")

	a.Initialize()

	go a.refreshTopics()
	a.Run(httpPort)
}
