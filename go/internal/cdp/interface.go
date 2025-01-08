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

	GetAllocatorType(context.Context, string) string
}

type IContextFactory interface {
	NewContext(context.Context) (context.Context, *chromedp.Context, context.CancelFunc)
}

type ISession interface {
	Run(context.Context, chromedp.Tasks) (context.Context, *chromedp.Context, error)
	Stop(context.Context)
}
