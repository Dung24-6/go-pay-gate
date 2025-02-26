package momo

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Dung24-6/go-pay-gate/pkg/gateway"
)

type MomoGateway struct {
	Config *gateway.Config
}

func NewMomoGateway(cfg *gateway.Config) *MomoGateway {
	return &MomoGateway{Config: cfg}
}

func (m *MomoGateway) CreatePayment(ctx context.Context, req *gateway.PaymentRequest) (*gateway.PaymentResponse, error) {
	params := map[string]string{
		"partnerCode": m.Config.MerchantID,
		"accessKey":   m.Config.ApiKey,
		"requestId":   req.OrderID,
		"amount":      fmt.Sprintf("%.0f", req.Amount),
		"orderId":     req.OrderID,
		"orderInfo":   req.Description,
		"returnUrl":   req.RedirectURL,
		"notifyUrl":   req.WebhookURL,
		"extraData":   "",
		"requestType": "captureMoMoWallet",
	}

	rawSignature := fmt.Sprintf("accessKey=%s&amount=%s&extraData=%s&notifyUrl=%s&orderId=%s&orderInfo=%s&partnerCode=%s&requestId=%s&requestType=%s",
		params["accessKey"],
		params["amount"],
		params["extraData"],
		params["notifyUrl"],
		params["orderId"],
		params["orderInfo"],
		params["partnerCode"],
		params["requestId"],
		params["requestType"],
	)
	h := hmac.New(sha256.New, []byte(m.Config.ApiSecret))
	h.Write([]byte(rawSignature))
	signature := hex.EncodeToString(h.Sum(nil))

	paymentURL := fmt.Sprintf("%s?partnerCode=%s&accessKey=%s&requestId=%s&orderId=%s&signature=%s",
		m.Config.ApiEndpoint,
		params["partnerCode"],
		params["accessKey"],
		params["requestId"],
		params["orderId"],
		signature,
	)

	resp := &gateway.PaymentResponse{
		TransactionID: req.OrderID,
		PaymentURL:    paymentURL,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Status:        gateway.StatusPending,
		OrderID:       req.OrderID,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(req.ExpiryDuration),
		PaymentMethod: "momo",
		RawResponse: map[string]interface{}{
			"params":    params,
			"signature": signature,
		},
	}
	return resp, nil
}

func (m *MomoGateway) QueryStatus(ctx context.Context, id string) (*gateway.PaymentStatus, error) {
	return &gateway.PaymentStatus{
		TransactionID:  id,
		OrderID:        id,
		Status:         gateway.StatusPending,
		Amount:         0,
		Currency:       "VND",
		PaymentMethod:  "momo",
		PaidAmount:     0,
		RefundedAmount: 0,
		PaidAt:         time.Time{},
		LastUpdated:    time.Now(),
		RawStatus:      map[string]interface{}{"demo": "pending"},
	}, nil
}

func (m *MomoGateway) ProcessCallback(data []byte) (*gateway.CallbackResponse, error) {
	var callbackData map[string]interface{}
	if err := json.Unmarshal(data, &callbackData); err != nil {
		return nil, fmt.Errorf("invalid callback data: %w", err)
	}
	return &gateway.CallbackResponse{
		TransactionID:  callbackData["orderId"].(string),
		OrderID:        callbackData["orderId"].(string),
		Status:         fmt.Sprintf("%v", callbackData["resultCode"]),
		Amount:         0,
		Currency:       "VND",
		PaymentMethod:  "momo",
		PaidAmount:     0,
		PaidAt:         time.Now(),
		SignatureValid: true,
		RawData:        callbackData,
	}, nil
}
