package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type PageData struct {
	Message   string
	TodayDate time.Time
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}
	defer dbpool.Close()

	r := gin.Default()

	tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head><title>Hello</title></head>
		<body>
			<h1>{{.Message}}</h1>
			<p>Today's date from PostgreSQL- hi rahul {{.TodayDate}}</p>
		</body>
		</html>`))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) {
		var date time.Time
		err := dbpool.QueryRow(ctx, "SELECT NOW();").Scan(&date)
		if err != nil {
			c.String(http.StatusInternalServerError, "DB Error: %v", err)
			return
		}
		date = time.Now()
		c.HTML(http.StatusOK, "index", PageData{
			Message:   "Hello, World",
			TodayDate: date,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
