package rest

import (
	"github.com/Dung24-6/go-pay-gate/internal/store/aws"
	"github.com/Dung24-6/go-pay-gate/internal/store/kafka"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HandlerConfig struct {
	DB    *gorm.DB
	Redis *redis.Client
	AWS   *aws.AWSClients
	Kafka *kafka.KafkaClient
}

type Handler struct {
	db    *gorm.DB
	redis *redis.Client
	aws   *aws.AWSClients
	kafka *kafka.KafkaClient
}

func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{
		db:    cfg.DB,
		redis: cfg.Redis,
		aws:   cfg.AWS,
		kafka: cfg.Kafka,
	}
}

// SetupRoutes configures the HTTP routes
func SetupRoutes(r *gin.Engine, h *Handler) {
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
