package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/api"
	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/ambulance"
	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/seeder"
	"github.com/gin-contrib/cors"
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
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	engine.Use(corsMiddleware)
	// setup context update  middleware
	dbService := db_service.NewMongoService[ambulance.Patient](db_service.MongoServiceConfig{})
	defer dbService.Disconnect(context.Background())

	seedPtr := flag.Bool("seed", false, "Run the database seeder")
	flag.Parse()

	if *seedPtr {
		seeder.Seed(dbService)
	}

	engine.Use(func(ctx *gin.Context) {
		ctx.Set("db_service", dbService)
		ctx.Next()
	})
	// request routings
	handleFunctions := &ambulance.ApiHandleFunctions{
		PatientsAPI:  ambulance.PatientsApi(),
		IllnessesAPI: ambulance.IllnessesApi(),
	}
	ambulance.NewRouterWithGinEngine(engine, *handleFunctions)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
