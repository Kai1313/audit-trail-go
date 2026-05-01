package handlers

import (
	"auditservice/models"
	"auditservice/services"
	"net/http"

	"github.com/labstack/echo/v4" // Use v4
)

type AuditHandler struct {
	Service services.AuditService
}

// CreateLog handles the submission of a new audit log
// @Summary Create a new audit log
// @Description Receives logs from microservices and queues them for MySQL storage
// @Tags logs
// @Accept json
// @Produce json
// @Param log body models.AuditEntry true "Audit Log Data"
// @Success 202 {string} string "Accepted"
// @Failure 400 {object} echo.Map "Invalid format"
// @Router /logs [post]
func (h *AuditHandler) CreateLog(c echo.Context) error {
	var entry models.AuditEntry
	
	// Bind incoming JSON to the model
	if err := c.Bind(&entry); err != nil {
		// In v4, echo.Map is back and perfectly fine to use
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	// Fallback to Context IP if not provided in the payload
	if entry.IPAddress == "" {
		entry.IPAddress = c.RealIP()
	}

	// Send to the async service worker
	h.Service.ProcessLog(entry)

	// 202 Accepted is best practice for async/queued tasks
	return c.NoContent(http.StatusAccepted)
}