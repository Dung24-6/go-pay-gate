# Go Pay Gate

![GitHub Repo stars](https://img.shields.io/github/stars/Dung24-6/go-pay-gate?style=flat)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Dung24-6/go-pay-gate)

# Features

- Simple, consistent API for multiple payment providers
- Built-in support for popular payment gateways:
  - **VNPay**
  - **Momo**
  - **ZaloPay**
- Concurrent processing with goroutines
- Configurable retries and timeouts
- Transaction logging and monitoring
- Easy to extend for new payment providers

## Installation

```bash
go get github.com/Dung24-6/go-pay-gate
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    "github.com/Dung24-6/go-pay-gate/pkg/gateway"
)

func main() {
    // Initialize VNPay gateway
    vnpay := gateway.NewVNPay(&gateway.Config{
        MerchantID:  "your-merchant-id",
        ApiKey:      "your-api-key",
        Environment: gateway.EnvSandbox,
    })

    // Create payment request
    req := &gateway.PaymentRequest{
        Amount:      100000,
        OrderID:     "order-123",
        Description: "Test payment",
    }

    // Process payment
    resp, err := vnpay.CreatePayment(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Payment URL: %s", resp.PaymentURL)
}
```

## Configuration

Create `config.yaml`:

```yaml
server:
  port: ":8080"
  environment: "development"

database:
  host: "localhost"
  port: "3306"
  user: "root"
  password: "root"
  dbname: "gopaygate"
```

## Documentation

For detailed documentation and examples, please visit our documentation.

## Supported Payment Providers

- **VNPay**
- **Momo**
- **ZaloPay** *(coming soon)*

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Add some amazing feature'
   ```
4. Push to the branch:
   ```bash
   git push origin feature/amazing-feature
   ```
5. Open a Pull Request

## Development

### Requirements

- Go 1.22+
- Docker & Docker Compose
- MySQL
- Redis

### Setup development environment

```bash
# Clone repository
git clone https://github.com/Dung24-6/go-pay-gate.git

# Start dependencies
docker-compose up -d

# Run tests
go test ./...

# Run server
go run cmd/server/main.go
```

## Tech Stack
- **Go**
- **Gin**
- **GORM**
- **SQL**
- **Redis**
- **Docker**
- **Aws**
- **Kafka**



## Authors

- [@Dung24-6](https://github.com/Dung24-6)

## 🚀 About us
We are software engineer

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



