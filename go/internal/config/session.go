package config

import "github.com/playwright-community/playwright-go"

type SessionConfig struct {
	BrowserTypeLaunchOptions playwright.BrowserTypeLaunchOptions
	BrowserNewContextOptions playwright.BrowserNewContextOptions
	PageGotoOptions          playwright.PageGotoOptions
}
