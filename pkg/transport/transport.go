package transport

import "context"

type Transport interface {
	Start() error
	Stop(ctx context.Context) error
}
