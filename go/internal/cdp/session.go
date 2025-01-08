package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/config"
)

type Session struct {
	allocator        chromedp.Allocator
	allocatorFactory IAllocatorFactory
	browser          *chromedp.Browser
	cancelAllocator  context.CancelFunc
	cancelContext    context.CancelFunc
	cdpctx           *chromedp.Context
	cfg              config.SessionConfig
	url              string
}

func NewSession(
	ctx context.Context,
	allocatorFactory IAllocatorFactory,
	cfg config.SessionConfig,
	url string,
) (context.Context, *Session, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("NewSession")

	var err error

	result := &Session{
		allocatorFactory: allocatorFactory,
		cfg:              cfg,
		url:              url,
	}

	// passes the allocator via the context
	ctx, result.cdpctx, result.allocator, result.cancelAllocator = result.allocatorFactory.NewAllocator(
		ctx, result.url,
	)
	if err != nil {
		return ctx, nil, err
	}

	return ctx, result, nil
}

func (s *Session) Run(
	ctx context.Context,
	tasks chromedp.Tasks,
) (context.Context, *chromedp.Context, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("SendTask")

	var err error

	err = chromedp.Run(ctx, tasks)
	if err != nil {
		return ctx, nil, err
	}

	return ctx, s.cdpctx, nil
}

func (s *Session) Stop(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Stop")

	s.allocator.Wait()
	logger.Info().Msg("Stop: allocator done")

	select {
	case <-ctx.Done():
		logger.Info().Msg("Stop: context done")
		// cancel the context using the cancel function
		s.cancelContext()
		// cancel the allocator using the cancel function
		s.cancelAllocator()
	}
}
