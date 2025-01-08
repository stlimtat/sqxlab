package cdp

import (
	"context"
	reflect "reflect"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stlimtat/sqxlab/go/internal/urls"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlDiscovery := urls.NewMockIUrlDiscovery(ctrl)
			urlDiscovery.EXPECT().Discover(ctx).Return(tt.url, nil)
			factory := NewDefaultAllocatorFactory(ctx, tt.cfg, urlDiscovery)

			assert.Equal(t, tt.wantType, factory.GetAllocatorType(ctx))
			ctx, cdpctx, got, cancel := factory.NewAllocator(ctx, tt.url)
			assert.NotNil(t, ctx)
			assert.NotNil(t, cdpctx)
			assert.NotNil(t, got)
			assert.NotNil(t, cancel)

			assert.IsType(t, tt.wantClass, reflect.TypeOf(got))
		})
	}
}
