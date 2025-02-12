package rest

import (
	"github.com/gin-gonic/gin"
)

// Handler holds the dependencies for the HTTP handlers
type Handler struct {
	// Add your services here
}

// NewHandler creates a new HTTP handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes configures the HTTP routes
func SetupRoutes(r *gin.Engine) {
	h := NewHandler()

	// Health check
	r.GET("/health", h.HealthCheck)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		payments := v1.Group("/payments")
		{
			payments.POST("/create", h.CreatePayment)
			payments.GET("/status/:id", h.GetPaymentStatus)
			payments.POST("/callback", h.HandleCallback)
		}
	}
}

// Handler methods
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (h *Handler) CreatePayment(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, gin.H{"message": "not implemented"})
}

func (h *Handler) GetPaymentStatus(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, gin.H{"message": "not implemented"})
}

func (h *Handler) HandleCallback(c *gin.Context) {
	// TODO: Implement
	c.JSON(200, gin.H{"message": "not implemented"})
}
