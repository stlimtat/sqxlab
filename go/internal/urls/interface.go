package urls

import "context"

//go:generate mockgen -destination=mock.go -package=urls -source=interface.go
type IUrlDiscovery interface {
	Discover(ctx context.Context) (string, error)
}
