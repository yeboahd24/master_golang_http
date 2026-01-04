package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	client        *http.Client
	downstreamURL string
}

func NewService(client *http.Client, downstreamURL string) *Service {
	return &Service{
		client:        client,
		downstreamURL: downstreamURL,
	}
}

type (
	requestIDKeyType struct{}
	userIDKeyType    struct{}
)

var (
	requestIDKey = requestIDKeyType{}
	userIDKey    = userIDKeyType{}
)

type spanIDKeyType struct{}

var spanIDKey = spanIDKeyType{}

func WithRequestInfo(parent context.Context, requestID, userID string) context.Context {
	ctx := context.WithValue(parent, requestIDKey, requestID)
	ctx = context.WithValue(ctx, userIDKey, userID)
	return ctx
}

type Data struct {
	TransactionID string  `json:"transaction_id"`
	UserID        string  `json:"user_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

// Progation
func (s *Service) CallDownstream(ctx context.Context, data Data) error {
	spanID, _ := ctx.Value(spanIDKey).(string)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		s.downstreamURL,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if spanID != "" {
		req.Header.Set("X-Span-ID", spanID)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("downstream error: %s", resp.Status)
	}

	return nil
}
