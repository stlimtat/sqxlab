package cdp

import (
	"context"
	"testing"

	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	var tests = []struct {
		name string
		cfg  config.SessionConfig
	}{
		{
			name: "default",
			cfg:  config.SessionConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			cf := NewDefaultContextFactory(ctx, tt.cfg)
			ctx, cdpctx, cancel := cf.NewContext(ctx)
			assert.NotNil(t, ctx)
			assert.NotNil(t, cdpctx)
			assert.NotNil(t, cancel)
		})
	}
}
