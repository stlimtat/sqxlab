package config

import "github.com/chromedp/chromedp"

type SessionConfig struct {
	// URL is the websocket URL to connect to the Chrome DevTools Protocol.
	URL string `json:"url" mapstructure:"url"`

	// RemoteAllocatorOptions are the options to pass to the remote allocator.
	RemoteAllocatorOptions []chromedp.RemoteAllocatorOption `json:"remote_allocator_options" mapstructure:"remote_allocator_options"`

	// ContextOptions are the options to pass to the context.
	ContextOptions []chromedp.ContextOption `json:"context_options" mapstructure:"context_options"`
}
