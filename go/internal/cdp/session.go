package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/config"
)

type Session struct {
	allocator       chromedp.Allocator
	browser         *chromedp.Browser
	cancelAllocator context.CancelFunc
	cancelContext   context.CancelFunc
	cdpctx          *chromedp.Context
	cfg             config.SessionConfig
}

func NewSession(
	ctx context.Context,
	cfg config.SessionConfig,
) (*Session, context.Context, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("NewSession")

	var err error

	result := &Session{
		cfg: cfg,
	}

	// passes the allocator via the context
	ctx, result.cancelAllocator = chromedp.NewRemoteAllocator(
		ctx,
		cfg.URL,
		result.cfg.RemoteAllocatorOptions...,
	)
	if err != nil {
		return nil, nil, err
	}

	// passes the chromedp context via the context
	ctx, result.cancelContext = chromedp.NewContext(ctx, result.cfg.ContextOptions...)

	result.cdpctx = chromedp.FromContext(ctx)
	result.allocator = result.cdpctx.Allocator

	return result, ctx, nil
}

func (s *Session) SendTask(
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
