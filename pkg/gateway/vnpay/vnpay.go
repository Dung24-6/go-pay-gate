package vnpay

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/Dung24-6/go-pay-gate/internal/logging"
	"github.com/Dung24-6/go-pay-gate/pkg/gateway"
	"go.uber.org/zap"
)

type VNPayGateway struct {
	Config *gateway.Config
}

func NewVNPayGateway(cfg *gateway.Config) *VNPayGateway {
	return &VNPayGateway{
		Config: cfg,
	}
}

func (v *VNPayGateway) CreatePayment(ctx context.Context, req *gateway.PaymentRequest) (*gateway.PaymentResponse, error) {
	params := map[string]string{
		"vnp_Version":    "2.1.0",
		"vnp_Command":    "pay",
		"vnp_TmnCode":    v.Config.MerchantID,
		"vnp_Amount":     fmt.Sprintf("%.0f", req.Amount*100),
		"vnp_CurrCode":   req.Currency,
		"vnp_TxnRef":     req.OrderID,
		"vnp_OrderInfo":  req.Description,
		"vnp_OrderType":  "other",
		"vnp_Locale":     "vn",
		"vnp_ReturnUrl":  req.RedirectURL,
		"vnp_IpAddr":     "127.0.0.1",
		"vnp_CreateDate": time.Now().Format("20060102150405"),
		"vnp_ExpireDate": time.Now().Add(req.ExpiryDuration).Format("20060102150405"),
	}

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var queryStrings []string
	var hashData strings.Builder
	for _, k := range keys {
		value := params[k]
		if len(value) > 0 {
			queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", k, url.QueryEscape(value)))
			hashData.WriteString(fmt.Sprintf("%s=%s&", k, value))
		}
	}
	hashDataStr := strings.TrimRight(hashData.String(), "&")

	h := md5.New()
	h.Write([]byte(v.Config.ApiSecret + hashDataStr))
	secureHash := hex.EncodeToString(h.Sum(nil))

	queryString := strings.Join(queryStrings, "&")
	paymentURL := fmt.Sprintf("%s?%s&vnp_SecureHash=%s", v.Config.ApiEndpoint, queryString, secureHash)

	resp := &gateway.PaymentResponse{
		TransactionID: req.OrderID,
		PaymentURL:    paymentURL,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Status:        gateway.StatusPending,
		OrderID:       req.OrderID,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(req.ExpiryDuration),
		PaymentMethod: "vnpay",
		RawResponse: map[string]interface{}{
			"params": params,
		},
	}
	logging.Logger.Info("VNPay payment created successfully",
		zap.String("order_id", req.OrderID),
	)
	return resp, nil
}

func (v *VNPayGateway) QueryStatus(ctx context.Context, id string) (*gateway.PaymentStatus, error) {
	status := &gateway.PaymentStatus{
		TransactionID:  id,
		OrderID:        id,
		Status:         gateway.StatusPending,
		Amount:         0,
		Currency:       "VND",
		PaymentMethod:  "vnpay",
		PaidAmount:     0,
		RefundedAmount: 0,
		PaidAt:         time.Time{},
		LastUpdated:    time.Now(),
		RawStatus:      map[string]interface{}{"simulated": true},
	}
	return status, nil
}

func (v *VNPayGateway) ProcessCallback(data []byte) (*gateway.CallbackResponse, error) {
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return nil, fmt.Errorf("invalid callback data: %w", err)
	}

	receivedHash := values.Get("vnp_SecureHash")
	if receivedHash == "" {
		return nil, fmt.Errorf("missing secure hash")
	}
	values.Del("vnp_SecureHash")

	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var hashData strings.Builder
	for _, k := range keys {
		value := values.Get(k)
		if len(value) > 0 {
			hashData.WriteString(fmt.Sprintf("%s=%s&", k, value))
		}
	}
	hashDataStr := strings.TrimRight(hashData.String(), "&")

	h := md5.New()
	h.Write([]byte(v.Config.ApiSecret + hashDataStr))
	calculatedHash := hex.EncodeToString(h.Sum(nil))
	signatureValid := (calculatedHash == receivedHash)

	callbackResp := &gateway.CallbackResponse{
		TransactionID:  values.Get("vnp_TxnRef"),
		OrderID:        values.Get("vnp_TxnRef"),
		Status:         values.Get("vnp_TransactionStatus"),
		Amount:         0,
		Currency:       values.Get("vnp_CurrCode"),
		PaymentMethod:  "vnpay",
		PaidAmount:     0,
		PaidAt:         time.Now(),
		SignatureValid: signatureValid,
		RawData:        map[string]interface{}{},
	}
	for key, val := range values {
		if len(val) > 0 {
			callbackResp.RawData[key] = val[0]
		}
	}
	return callbackResp, nil
}
