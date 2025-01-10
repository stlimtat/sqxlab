package config

import "github.com/chromedp/chromedp"

type SessionConfig struct {
	// ContextOptions are the options to pass to the context.
	ContextOptions []chromedp.ContextOption `json:"context_options" mapstructure:"context_options"`

	// ExecAllocatorOptions are the options to pass to the exec allocator.
	ExecAllocatorOptions []chromedp.ExecAllocatorOption `json:"exec_allocator_options" mapstructure:"exec_allocator_options"`

	// RemoteAllocatorOptions are the options to pass to the remote allocator.
	RemoteAllocatorOptions []chromedp.RemoteAllocatorOption `json:"remote_allocator_options" mapstructure:"remote_allocator_options"`
}

type SessionContext struct {
	SessionConfig
	// URL for the current session
	URL string `json:"url" mapstructure:"url"`
}
