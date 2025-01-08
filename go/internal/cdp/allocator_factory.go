package cdp

import (
	"context"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/config"
)

const (
	AllocatorTypeDefault = "default"
	AllocatorTypeRemote  = "remote"
)

type DefaultAllocatorFactory struct {
	allocator_type string
	cfg            config.SessionConfig
	url            string
}

func NewDefaultAllocatorFactory(
	ctx context.Context,
	cfg config.SessionConfig,
	url string,
) *DefaultAllocatorFactory {
	result := &DefaultAllocatorFactory{
		cfg: cfg,
		url: url,
	}

	return result
}

func (a *DefaultAllocatorFactory) NewAllocator(
	ctx context.Context,
	url string,
) (
	context.Context,
	*chromedp.Context,
	chromedp.Allocator,
	context.CancelFunc,
) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("NewAllocator")

	a.allocator_type = a.GetAllocatorType()

	var result chromedp.Allocator
	var cancelFunc context.CancelFunc
	switch a.allocator_type {
	case AllocatorTypeRemote:
		ctx, cancelFunc = chromedp.NewRemoteAllocator(
			ctx, a.url,
			a.cfg.RemoteAllocatorOptions...,
		)
	case AllocatorTypeDefault:
		ctx, cancelFunc = chromedp.NewExecAllocator(
			ctx, a.cfg.ExecAllocatorOptions...,
		)
	default:
		ctx, cancelFunc = chromedp.NewExecAllocator(
			ctx, a.cfg.ExecAllocatorOptions...,
		)
	}

	cdpctx := chromedp.FromContext(ctx)
	result = cdpctx.Allocator

	return ctx, cdpctx, result, cancelFunc
}

func (a *DefaultAllocatorFactory) GetAllocatorType() string {
	if a.url != "" && strings.HasPrefix(a.url, "ws://") {
		return AllocatorTypeRemote
	}
	return AllocatorTypeDefault
}
