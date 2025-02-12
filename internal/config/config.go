package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	AWS      AWSConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
	Port         string
	GRPCPort     string
	Environment  string
	ReadTimeout  time.Duration // Timeout for reading request body
	WriteTimeout time.Duration // Timeout for writing response
	IdleTimeout  time.Duration // Idle connection timeout
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int           // Maximum number of open connections
	MaxIdleConns    int           // Maximum number of idle connections
	ConnMaxLifetime time.Duration // Maximum amount of time a connection may be reused
}

type RedisConfig struct {
	Host         string
	Port         string
	Password     string
	DB           int
	PoolSize     int           // Maximum number of socket connections
	MinIdleConns int           // Minimum number of idle connections
	MaxRetries   int           // Maximum number of retries
	IdleTimeout  time.Duration // Close connections after remaining idle for this duration
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	S3Bucket        string
	SQSQueueURL     string
	SNSTopicARN     string
}

type KafkaConfig struct {
	Brokers         []string
	Topic           string
	ConsumerGroup   string
	MaxMessageBytes int           // Maximum message bytes that the consumer will fetch
	ReadTimeout     time.Duration // Timeout for reading messages
	WriteTimeout    time.Duration // Timeout for writing messages
}

// Load reads configuration from file or environment
func Load() (*Config, error) {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
