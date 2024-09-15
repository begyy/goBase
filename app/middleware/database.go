package middleware

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://begyy:@localhost:5432/golang")
	if err != nil {
		log.Fatal("Database connection failed:", err)
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}

func DBMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
