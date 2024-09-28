package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"goBase/app/middleware"
	"goBase/app/router"
	"log"
)

func main() {
	e := echo.New()
	db, err := middleware.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
	e.Use(middleware.DBMiddleware(db))
	// Routers
	router.ConfigureUserRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
