package urls

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog"

	"github.com/stlimtat/sqxlab/go/internal/config"
)

type DefaultUrlDiscovery struct {
	cfg config.UrlDiscoveryConfig
}

func NewDefaultUrlDiscovery(
	ctx context.Context,
	cfg config.UrlDiscoveryConfig,
) *DefaultUrlDiscovery {
	result := &DefaultUrlDiscovery{
		cfg: cfg,
	}

	return result
}

func (u *DefaultUrlDiscovery) Discover(ctx context.Context) (string, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("UrlDiscovery.Discover")
	var randomIndex int
	randomIndex = 0

	// discover url from a list
	if len(u.cfg.URLs) > 0 {
		// random select from the list
		randInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(u.cfg.URLs))))
		if err != nil {
			return "", err
		}
		randomIndex = int(randInt.Int64())
	}
	return u.cfg.URLs[randomIndex], nil
}
