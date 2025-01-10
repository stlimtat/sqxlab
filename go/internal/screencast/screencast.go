package screencast

import (
	"context"
	"reflect"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/cdp"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stlimtat/sqxlab/go/internal/urls"
)

type ScreencastFactory struct {
	AllocatorFactory cdp.IAllocatorFactory
	Cfg              config.ScreencastConfig
	Sessions         []cdp.ISession
	UrlDiscovery     urls.IUrlDiscovery
}

func NewScreencastFactory(
	ctx context.Context,
	allocatorFactory cdp.IAllocatorFactory,
	cfg config.ScreencastConfig,
	urlDiscovery urls.IUrlDiscovery,
) *ScreencastFactory {
	result := &ScreencastFactory{
		AllocatorFactory: allocatorFactory,
		Cfg:              cfg,
		UrlDiscovery:     urlDiscovery,
	}

	return result
}

func (s *ScreencastFactory) Run(
	ctx context.Context,
) error {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("ScreencastFactory.Run")

	url, err := s.UrlDiscovery.Discover(ctx)
	if err != nil {
		return err
	}

	ctx, session, err := cdp.NewSession(ctx, s.AllocatorFactory, s.Cfg.Session, url)
	if err != nil {
		return err
	}

	var body string
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(10 * time.Second),
		chromedp.OuterHTML("html", &body),
		page.StartScreencast().
			WithEveryNthFrame(5).
			WithFormat(page.ScreencastFormatPng).
			WithMaxHeight(768).
			WithMaxWidth(1024).
			WithQuality(50),
		chromedp.Sleep(10 * time.Second),
		page.StopScreencast(),
	}

	for idx, action := range tasks {
		sublogger := logger.With().
			Int("idx", idx).
			Str("action", reflect.TypeOf(action).String()).
			Logger()
		sublogger.Info().Msg("action")
		subCtx := sublogger.WithContext(ctx)
		_, _, err = session.Run(subCtx, chromedp.Tasks{action})
		if err != nil {
			sublogger.Error().Err(err).Msg("action failed")
			return err
		}
		sublogger.Info().Msg("action done")
	}

	return nil
}
