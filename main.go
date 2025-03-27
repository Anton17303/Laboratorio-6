package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/anton17/la-liga-tracker/internal/database"
	"github.com/anton17/la-liga-tracker/internal/handlers"
)

func main() {
	// Inicializar base de datos
	database.InitDatabase()

	// Configurar router
	router := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Grupo de rutas API
	api := router.Group("/api")
	{
		// CRUD de partidos
		api.GET("/matches", handlers.GetAllMatches)
		api.GET("/matches/:id", handlers.GetMatchByID)
		api.POST("/matches", handlers.CreateMatch)
		api.PUT("/matches/:id", handlers.UpdateMatch)
		api.DELETE("/matches/:id", handlers.DeleteMatch)

		// Rutas de eventos adicionales
		api.PATCH("/matches/:id/goals", handlers.RegisterGoal)
		api.PATCH("/matches/:id/yellowcards", handlers.RegisterYellowCard)
		api.PATCH("/matches/:id/redcards", handlers.RegisterRedCard)
		api.PATCH("/matches/:id/extratime", handlers.SetExtraTime)
	}

	// Puerto configurable via variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando en puerto %s", port)
	log.Fatal(router.Run(":" + port))
}