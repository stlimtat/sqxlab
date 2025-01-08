package cdp

import (
	"context"
	"fmt"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSessionRun(t *testing.T) {
	var tests = []struct {
		name          string
		cfg           config.SessionConfig
		url           string
		wantAllocator bool
		wantContext   *chromedp.Context
		wantCancel    bool
		wantErr       error
	}{
		{
			name:          "default",
			cfg:           config.SessionConfig{},
			url:           "ws://127.0.0.1:9222/",
			wantAllocator: true,
			wantContext:   nil,
			wantCancel:    true,
			wantErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAllocator := NewMockAllocator(ctrl)

			mockAllocatorFactory := NewMockIAllocatorFactory(ctrl)
			mockAllocatorFactory.EXPECT().NewAllocator(
				gomock.Any(), gomock.Any(),
			).DoAndReturn(func(
				ctx context.Context, url string,
			) (
				context.Context,
				*chromedp.Context,
				chromedp.Allocator,
				context.CancelFunc,
			) {
				var cancel context.CancelFunc
				if tt.wantCancel {
					cancel = func() { fmt.Println("cancel") }
				}
				return ctx, tt.wantContext, mockAllocator, cancel
			})

			ctx, got, err := NewSession(
				ctx,
				mockAllocatorFactory,
				tt.cfg,
				tt.url,
			)
			assert.NotNil(t, ctx)
			assert.NotNil(t, got)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantAllocator, got.allocator != nil)
			assert.Equal(t, tt.wantContext, got.cdpctx)
			assert.Equal(t, tt.wantCancel, got.cancelAllocator != nil)
		})
	}
}
