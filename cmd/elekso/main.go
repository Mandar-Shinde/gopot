package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"

	
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
var (
    repeat int
    db     *sql.DB = nil
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	//-----------------------------------------
	
	 router.GET("/test", func(c *gin.Context) {
        c.String(http.StatusOK, string("google mandar ... shinde "))
    })

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	router.Run(":" + port)
}