package main

import (
	"fmt"
	"log"

	"github.com/Set-Forget/Hackaton3/ms-catalog/config"
	"github.com/Set-Forget/Hackaton3/ms-catalog/handlers"
	"github.com/Set-Forget/Hackaton3/ms-catalog/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
    
    // Get Supabase connection details from environment variables
    host := "aws-0-us-east-1.pooler.supabase.com"
    port := "6543"
    user := "postgres.hlflpxnttokwivejeffn"
    password := "hackaton3setandforget"
    dbname := "postgres"

    if host == "" || port == "" || user == "" || password == "" || dbname == "" {
        fmt.Println(host)
        fmt.Println(port)
        fmt.Println(user)
        fmt.Println(dbname)
        log.Fatal("Missing required environment variables for database connection")
    }

    // Initialize database connection
    db := config.InitDB(host, port, user, password, dbname)
    defer db.Close()

    // Initialize repositories and handlers
    productRepo := repositories.NewProductRepository(db)
    catalogHandler := handlers.NewCatalogHandler(productRepo)

    // Setup Fiber app
    app := fiber.New()
    api := app.Group("/api")
    
    // Routes
    api.Get("/catalog", catalogHandler.GetCatalog)

    // Start server
    log.Fatal(app.Listen(":3000"))
}