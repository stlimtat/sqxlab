package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
)

//go:generate mockgen -destination=chromedp_mock.go -package=cdp github.com/chromedp/chromedp Allocator
//go:generate mockgen -destination=mock.go -package=cdp github.com/stlimtat/sqxlab/go/internal/cdp IAllocatorFactory,IContextFactory,ISession
type IAllocatorFactory interface {
	NewAllocator(
		context.Context, string,
	) (context.Context, *chromedp.Context, chromedp.Allocator, context.CancelFunc)

	GetAllocatorType() string
}

type IContextFactory interface {
	NewContext(ctx context.Context) (context.Context, *chromedp.Context, context.CancelFunc)
}

type ISession interface {
	Run(ctx context.Context, tasks chromedp.Tasks) (context.Context, *chromedp.Context, error)
	Stop(ctx context.Context)
}
