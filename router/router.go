package router

import (
	"auditservice/handlers"
	_ "auditservice/docs" // Load generated swagger docs

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func SetupRouter(e *echo.Echo, h *handlers.AuditHandler) {
	// 1. Global Middleware
	// In v4, these are functions: you MUST include the ()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 2. Swagger Documentation
	// Native compatibility: no wrapper needed in v4
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 3. Health Check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "UP"})
	})

	// 4. API Group
	api := e.Group("/api/v1")

	// 5. Audit Log Routes
	api.POST("/logs", h.CreateLog)
}