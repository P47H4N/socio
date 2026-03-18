package main

import (
	"log"

	"github.com/P47H4N/socio/cmd"
	"github.com/P47H4N/socio/internals/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	
	config, err := cmd.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config.", err)
	}

	db, err := database.NewDB(config)
	if err != nil {
		log.Fatalln("Unable to connect with database.", err)
	}

	route := router.Group("/api/v1")

	cmd.Start(route, db)

	router.Run()
}