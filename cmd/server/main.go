package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dung24-6/go-pay-gate/api/rest"
	"github.com/Dung24-6/go-pay-gate/internal/config"
	"github.com/Dung24-6/go-pay-gate/internal/logging"
	"github.com/Dung24-6/go-pay-gate/internal/repository"
	"github.com/Dung24-6/go-pay-gate/internal/store/aws"
	"github.com/Dung24-6/go-pay-gate/internal/store/kafka"
	"github.com/Dung24-6/go-pay-gate/internal/store/mysql"
	"github.com/Dung24-6/go-pay-gate/internal/store/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	redis       *redis.Client
	aws         *aws.AWSClients
	kafka       *kafka.KafkaClient
	httpServer  *http.Server
	paymentRepo repository.PaymentRepository
}

func main() {
	// Initial logger
	logging.InitLogger()
	defer logging.Sync()

	logging.Logger.Info("Starting go-pay-gate server...")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := &App{}

	// Initialize DB
	app.db, err = mysql.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	app.paymentRepo = repository.NewMySQLPaymentRepository(app.db)

	// Initialize Redis
	app.redis, err = redis.NewRedisClient(&cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Initialize AWS clients
	app.aws, err = aws.NewAWSClients(&cfg.AWS)
	if err != nil {
		log.Fatalf("Failed to initialize AWS clients: %v", err)
	}

	// Initialize Kafka
	app.kafka, err = kafka.NewKafkaClient(&cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka client: %v", err)
	}
	defer app.kafka.Writer.Close()
	defer app.kafka.Reader.Close()

	// Setup Gin
	r := gin.Default()

	// Initialize handler with dependencies
	handler := rest.NewHandler(rest.HandlerConfig{
		DB:    app.db,
		Redis: app.redis.Client,
		AWS:   app.aws,
		Kafka: app.kafka,
	})

	// Setup routes
	rest.SetupRoutes(r, handler)

	// Create server
	app.httpServer = &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	// Start server in goroutine
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Start Kafka consumer in goroutine
	go app.consumeKafkaMessages()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func (app *App) consumeKafkaMessages() {
	for {
		message, err := app.kafka.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
			continue
		}
		// Process message
		log.Printf("Received message: %s", string(message.Value))
	}
}

func (app *App) shutdown(ctx context.Context) error {
	// Shutdown HTTP server
	if err := app.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	// Close database connection
	sqlDB, err := app.db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}

	// Close Redis connection
	if err := app.redis.Close(); err != nil {
		return err
	}

	return nil
}
