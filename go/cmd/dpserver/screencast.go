/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/stlimtat/sqxlab/go/internal/cdp"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stlimtat/sqxlab/go/internal/telemetry"
	"github.com/stlimtat/sqxlab/go/internal/urls"
	"golang.org/x/sync/errgroup"
)

type screencastCmd struct {
	cmd        *cobra.Command
	screencast *Screencast
}

func newScreencastCmd(ctx context.Context) (*screencastCmd, *cobra.Command) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("newScreencastCmd")
	var err error

	result := &screencastCmd{}

	// serverCmd represents the server command
	result.cmd = &cobra.Command{
		Use:   "screencast",
		Short: "Run screencast on a single remote shell",
		Long:  `Evoke the Screencast via Chrome Devtools Protocol via a single headless shell`,
		Args: func(_ *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			result.screencast = newScreencast(cmd, args)
			err = result.screencast.Run(cmd.Context())
			return err
		},
	}

	return result, result.cmd
}

type Screencast struct {
	AllocatorFactory cdp.IAllocatorFactory
	Cfg              config.ScreencastConfig
	UrlDiscovery     urls.IUrlDiscovery
}

func newScreencast(
	cmd *cobra.Command,
	_ []string,
) *Screencast {
	ctx := cmd.Context()
	result := &Screencast{}

	result.Cfg = config.NewScreencastConfig(ctx)

	if result.Cfg.Debug {
		telemetry.SetGlobalLogLevel(zerolog.DebugLevel)
	}

	result.AllocatorFactory = cdp.NewDefaultAllocatorFactory(ctx, result.Cfg.Session)
	result.UrlDiscovery = urls.NewDefaultUrlDiscovery(ctx, result.Cfg.UrlDiscovery)

	return result
}

func (s *Screencast) Run(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Screencast.Run")
	eg, ctx := errgroup.WithContext(ctx)
	var err error

	eg.Go(func() error {
		return nil
	})

	err = eg.Wait()
	if err != nil {
		logger.Error().Err(err).Msg("errgroup Wait")
	}
	return err
}
