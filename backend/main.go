package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

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

	dbType := utils.GetEnv("DB_TYPE", "sqlite3")

	httpPort := utils.GetEnv("HTTP_PORT", "8080")
	sqlitePath := utils.GetEnv("SQLITE_DB_PATH", "rfd.db")

	// Determine database
	switch {
	case dbType == "sqlite3":
		log.Debug().Msgf("Using sqlite3 (path: " + sqlitePath + ")")
		a.Initialize(dbType, sqlitePath)
	default:
		log.Fatal().Msgf("Unsupported database: " + dbType)
	}

	a.Run(httpPort)
}
