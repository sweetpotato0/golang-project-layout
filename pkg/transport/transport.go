package transport

import "context"

// Server .
type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
