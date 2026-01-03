package health

import (
	"context"
	"github.com/alimarzban99/ecommerce/pkg/cache"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct{}

func NewHealthHandler() *Handler {
	return &Handler{}
}

type Status struct {
	Status    string                   `json:"status"`
	Timestamp string                   `json:"timestamp"`
	Services  map[string]ServiceStatus `json:"services"`
}

type ServiceStatus struct {
	Status       string `json:"status"`
	Message      string `json:"message,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
}

func (h *Handler) Health(c *gin.Context) {
	status := Status{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.DateTime),
		Services:  make(map[string]ServiceStatus),
	}

	dbStatus := h.checkDatabase()
	status.Services["database"] = dbStatus

	redisStatus := h.checkRedis()
	status.Services["redis"] = redisStatus

	if dbStatus.Status != "healthy" || redisStatus.Status != "healthy" {
		status.Status = "unhealthy"
		c.JSON(503, gin.H{
			"success": false,
			"data":    status,
		})
		return
	}

	response.SuccessResponse(c, status)
}

func (h *Handler) checkDatabase() ServiceStatus {
	start := time.Now()

	db := database.DB()
	if db == nil {
		return ServiceStatus{
			Status:  "unhealthy",
			Message: "Database connection not initialized",
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return ServiceStatus{
			Status:  "unhealthy",
			Message: "Failed to get database instance: " + err.Error(),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return ServiceStatus{
			Status:       "unhealthy",
			Message:      "Database ping failed: " + err.Error(),
			ResponseTime: time.Since(start).String(),
		}
	}

	return ServiceStatus{
		Status:       "healthy",
		Message:      "Database connection is active",
		ResponseTime: time.Since(start).String(),
	}
}

func (h *Handler) checkRedis() ServiceStatus {
	start := time.Now()

	client, err := cache.Client()
	if err != nil {
		return ServiceStatus{
			Status:  "unhealthy",
			Message: "Redis client not initialized: " + err.Error(),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return ServiceStatus{
			Status:       "unhealthy",
			Message:      "Redis ping failed: " + err.Error(),
			ResponseTime: time.Since(start).String(),
		}
	}

	return ServiceStatus{
		Status:       "healthy",
		Message:      "Redis connection is active",
		ResponseTime: time.Since(start).String(),
	}
}
