package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/app/migration"
	"awesomeProject/internal/app/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3005
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	_ = config.ConnectDB()
	db := config.ConnectDB()
	migrationDB(db)

	init := config.Init()
	app := routes.Route(init)
	err := app.Run(":" + port)
	if err != nil {
		return
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	return
}

func migrationDB(db *gorm.DB) {
	defer config.CloseDB(db)
	migration.Migrate(db)
}
