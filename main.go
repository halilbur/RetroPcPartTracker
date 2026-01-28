package main

import (
	"database/sql"
	"log"
	"retroPcPartTracker/handlers"
	"retroPcPartTracker/store"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/pcparts?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize store
	partStore := store.NewPartStore(db)

	// Initialize handlers
	h := handlers.NewHandlers(partStore)

	// Setup Gin router
	router := gin.Default()

	// Serve static files (CSS)
	router.Static("/static", "./static")

	// Routes
	router.GET("/", h.HandleHome)
	router.GET("/search", h.HandleSearch)
	router.GET("/parts/:type", h.HandlePartsByType)
	router.GET("/part/:id", h.HandlePartDetail)

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
