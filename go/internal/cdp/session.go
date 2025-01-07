package cdp

import (
	"context"
	"fmt"

	playwright "github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog"
	"github.com/stlimtat/sqxlab/go/internal/config"
)

type Session struct {
	browser    playwright.Browser
	cdpsession playwright.CDPSession
	context    playwright.BrowserContext
	cfg        config.SessionConfig
	page       playwright.Page
	playwright *playwright.Playwright
}

func NewSession(
	ctx context.Context,
	cfg config.SessionConfig,
) *Session {
	result := &Session{
		cfg: cfg,
	}

	return result
}

func (s *Session) Start(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("session.start")
	var err error
	s.playwright, err = playwright.Run()
	if err != nil {
		return err
	}
	logger.Info().Msg("playwright.Run")
	logger.Info().Msg("playwright.Chromium.Launch")
	s.browser, err = s.playwright.Chromium.Launch(s.cfg.BrowserTypeLaunchOptions)
	if err != nil {
		return err
	}
	s.cdpsession, err = s.browser.NewBrowserCDPSession()
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) GetBrowser(ctx context.Context) playwright.Browser {
	return s.browser
}

func (s *Session) GetCDPSession(ctx context.Context) playwright.CDPSession {
	return s.cdpsession
}

func (s *Session) Send(ctx context.Context, method string, params map[string]interface{}) (interface{}, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("cdpsession.send")
	return s.cdpsession.Send(method, params)
}

func (s *Session) Goto(
	ctx context.Context,
	url string,
	cdpMethod string,
	cdpParams map[string]interface{},
) (playwright.Response, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("session.goto")
	var err error
	s.context, err = s.browser.NewContext(s.cfg.BrowserNewContextOptions)
	if err != nil {
		return nil, err
	}
	logger.Info().Msg("browser.NewContext")
	s.page, err = s.context.NewPage()
	if err != nil {
		return nil, err
	}
	logger.Info().Msg("context.NewPage")
	response, err := s.page.Goto(url, s.cfg.PageGotoOptions)
	if err != nil {
		return nil, err
	}
	logger.Info().Msg("page.Goto")
	if response.Status() != 200 {
		return nil, fmt.Errorf("status: %d", response.Status())
	}
	return response, nil
}

func (s *Session) Stop(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("session.stop")
	var err error
	if s.cdpsession != nil {
		err = s.cdpsession.Detach()
		if err != nil {
			return err
		}
	}
	if s.browser != nil {
		err = s.browser.Close(playwright.BrowserCloseOptions{
			Reason: playwright.String("session.stop"),
		})
		if err != nil {
			return err
		}
	}
	if s.playwright != nil {
		err = s.playwright.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
