package main

import (
	"auditservice/handlers"
	"auditservice/repositories"
	"auditservice/router"
	"auditservice/services"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4" // Confirmed v4
)

// @title Audit Trail Service API
// @version 1.0
// @description This service handles technical error logging and audit trails.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No .env file found, using system environment variables")
	}

	// 2. Database Connection Logic
	// Note: parseTime=true is often needed for MySQL to scan dates into Go time.Time
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Database unreachable: %v", err)
	}
	log.Println("✅ Database connected successfully")

	// 3. Dependency Injection (Layered Architecture)
	repo := repositories.NewAuditRepository(db)
	svc := services.NewAuditService(repo)
	hdl := &handlers.AuditHandler{Service: svc}

	// 4. Initialize Echo
	e := echo.New()

	// 5. Load Router Configuration
	router.SetupRouter(e, hdl)

	// 6. Server Start Configuration
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	fmt.Printf("\n🚀 Audit Service started successfully!")
	fmt.Printf("\n📡 Listening on port: %s", appPort)
	fmt.Printf("\n📖 Swagger UI: http://localhost:%s/swagger/index.html\n\n", appPort)

	// Start server
	if err := e.Start(":" + appPort); err != nil {
		e.Logger.Fatal("❌ Server failed to shut down: ", err)
	}
}