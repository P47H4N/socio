package main

import (
	"log"
	"time"

	"github.com/P47H4N/socio/cmd"
	"github.com/P47H4N/socio/internals/database"
	"github.com/P47H4N/socio/internals/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	config, err := cmd.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config.", err)
	}
	
	helpers.LoadJWT(config.JWTToken)

	db, err := database.NewDB(config)
	if err != nil {
		log.Fatalln("Unable to connect with database.", err)
	}

	route := router.Group("/api/v1")

	cmd.Start(route, db)

	router.Run()
}