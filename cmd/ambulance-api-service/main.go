package main

import (
	"log"
	"os"
	"strings"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/api"
	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/ambulance"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	// request routings
	handleFunctions := &ambulance.ApiHandleFunctions{
		PatientsAPI: ambulance.PatientsApi(),
	}
	ambulance.NewRouterWithGinEngine(engine, *handleFunctions)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
