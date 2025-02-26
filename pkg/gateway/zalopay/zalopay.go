package zalopay

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Dung24-6/go-pay-gate/pkg/gateway"
)

type ZaloPayGateway struct {
	Config *gateway.Config
}

func NewZaloPayGateway(cfg *gateway.Config) *ZaloPayGateway {
	return &ZaloPayGateway{Config: cfg}
}

func (z *ZaloPayGateway) CreatePayment(ctx context.Context, req *gateway.PaymentRequest) (*gateway.PaymentResponse, error) {
	params := map[string]string{
		"app_id":      z.Config.MerchantID,
		"app_user":    req.CustomerID,
		"amount":      fmt.Sprintf("%.0f", req.Amount),
		"app_time":    fmt.Sprintf("%d", time.Now().UnixNano()/1e6),
		"item":        req.Description,
		"embed_data":  "",
		"bank_code":   "",
		"description": req.Description,
		"order_id":    req.OrderID,
	}

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var rawData strings.Builder
	for _, k := range keys {
		rawData.WriteString(fmt.Sprintf("%s=%s&", k, params[k]))
	}
	signData := strings.TrimRight(rawData.String(), "&")
	h := hmac.New(sha256.New, []byte(z.Config.ApiSecret))
	h.Write([]byte(signData))
	signature := hex.EncodeToString(h.Sum(nil))

	paymentURL := fmt.Sprintf("%s?order_id=%s&signature=%s", z.Config.ApiEndpoint, req.OrderID, signature)
	resp := &gateway.PaymentResponse{
		TransactionID: req.OrderID,
		PaymentURL:    paymentURL,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Status:        gateway.StatusPending,
		OrderID:       req.OrderID,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(req.ExpiryDuration),
		PaymentMethod: "zalopay",
		RawResponse: map[string]interface{}{
			"params":    params,
			"signature": signature,
		},
	}
	return resp, nil
}

func (z *ZaloPayGateway) QueryStatus(ctx context.Context, id string) (*gateway.PaymentStatus, error) {
	return &gateway.PaymentStatus{
		TransactionID:  id,
		OrderID:        id,
		Status:         gateway.StatusPending,
		Amount:         0,
		Currency:       "VND",
		PaymentMethod:  "zalopay",
		PaidAmount:     0,
		RefundedAmount: 0,
		PaidAt:         time.Time{},
		LastUpdated:    time.Now(),
		RawStatus:      map[string]interface{}{"demo": "pending"},
	}, nil
}

func (z *ZaloPayGateway) ProcessCallback(data []byte) (*gateway.CallbackResponse, error) {
	var callbackData map[string]interface{}
	if err := json.Unmarshal(data, &callbackData); err != nil {
		return nil, fmt.Errorf("invalid callback data: %w", err)
	}
	return &gateway.CallbackResponse{
		TransactionID:  callbackData["order_id"].(string),
		OrderID:        callbackData["order_id"].(string),
		Status:         fmt.Sprintf("%v", callbackData["result_code"]),
		Amount:         0,
		Currency:       "VND",
		PaymentMethod:  "zalopay",
		PaidAmount:     0,
		PaidAt:         time.Now(),
		SignatureValid: true,
		RawData:        callbackData,
	}, nil
}
