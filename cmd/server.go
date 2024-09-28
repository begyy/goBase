package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"goBase/app/config"
	"goBase/app/middleware"
	"goBase/app/router"
	_ "goBase/docs"
	"log"
)

// @title goBase API
// @version 1.0
// @description GoBase
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1/
func main() {
	e := echo.New()
	config.Load()
	db, err := middleware.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
	e.Use(middleware.DBMiddleware(db))
	// Routers
	e.GET("/swagger/*", echoSwagger.WrapHandler) // Swagger
	router.ConfigureUserRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
