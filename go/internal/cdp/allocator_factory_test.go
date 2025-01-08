package cdp

import (
	"context"
	reflect "reflect"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAllocator(t *testing.T) {
	var tests = []struct {
		name      string
		cfg       config.SessionConfig
		url       string
		wantType  string
		wantClass reflect.Type
	}{
		{
			name:      "default",
			cfg:       config.SessionConfig{},
			url:       "ws://127.0.0.1:9222/",
			wantType:  AllocatorTypeRemote,
			wantClass: reflect.TypeOf(chromedp.RemoteAllocator{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			factory := NewDefaultAllocatorFactory(ctx, tt.cfg, tt.url)

			assert.Equal(t, tt.wantType, factory.GetAllocatorType())
			ctx, cdpctx, got, cancel := factory.NewAllocator(ctx, tt.url)
			assert.NotNil(t, ctx)
			assert.NotNil(t, cdpctx)
			assert.NotNil(t, got)
			assert.NotNil(t, cancel)

			assert.IsType(t, tt.wantClass, reflect.TypeOf(got))
		})
	}
}
