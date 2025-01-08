package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/config"
)

type DefaultContextFactory struct {
	cfg config.SessionConfig
}

func NewDefaultContextFactory(
	ctx context.Context,
	cfg config.SessionConfig,
) *DefaultContextFactory {
	result := &DefaultContextFactory{
		cfg: cfg,
	}

	return result
}

func (cf *DefaultContextFactory) NewContext(
	ctx context.Context,
) (context.Context, *chromedp.Context, context.CancelFunc) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("NewContext")

	ctx, cancel := chromedp.NewContext(ctx, cf.cfg.ContextOptions...)
	cdpctx := chromedp.FromContext(ctx)

	return ctx, cdpctx, cancel
}
