CREATE TABLE payments (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id        VARCHAR(255) NOT NULL UNIQUE,
    transaction_id  VARCHAR(255) NOT NULL UNIQUE,
    amount          DECIMAL(18,2) NOT NULL,
    currency        VARCHAR(10) NOT NULL DEFAULT 'VND',
    payment_method  VARCHAR(50) NOT NULL,
    status          VARCHAR(20) NOT NULL DEFAULT 'pending',
    customer_id     VARCHAR(255),
    customer_name   VARCHAR(255),
    customer_email  VARCHAR(255),
    customer_phone  VARCHAR(50),
    redirect_url    TEXT,
    webhook_url     TEXT,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    payment_id      BIGINT NOT NULL,
    status          VARCHAR(20) NOT NULL,
    amount          DECIMAL(18,2) NOT NULL,
    paid_at         TIMESTAMP NULL,
    error_code      VARCHAR(50),
    error_message   TEXT,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_id) REFERENCES payments(id) ON DELETE CASCADE
);
