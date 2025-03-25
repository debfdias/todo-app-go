package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (handler *HealthHandler) CheckHealth(c *gin.Context) {
	dbStatus := "OK"
	if err := handler.db.Exec("SELECT 1").Error; err != nil {
		dbStatus = "Unhealthy"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "🚀 Rocket go! API is running",
		"database": dbStatus,
		"time":     time.Now().UTC().Format(time.RFC3339),
	})
}