// Package webhook defines basic functionality for webhooks
package webhook

import "context"

// Webhook interface basic functionality
type Webhook interface {
	EnsureWebhook(ctx context.Context, name string, opts ...EnsureOption) (*HTTPWebhook, error)
	ReadEvent(ctx context.Context, webhookID string, outValue interface{}) (bool, error)
}

// HTTPWebhook model will hold webhook data
type HTTPWebhook struct {
	ID   string
	URL  string
	Name string
}

// EnsureOption func to help pass options when needed, and also evaluate the options
type EnsureOption func(*EnsureOptions)

// EnsureOptions contains webhook options when ensuring
type EnsureOptions struct {
	Read bool // Should we subscribe and read webhook events
}

// EnsureWithRead will set Read opt true
func EnsureWithRead() EnsureOption {
	return func(opt *EnsureOptions) { opt.Read = true }
}
