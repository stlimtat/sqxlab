/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	zerologgin "github.com/go-mods/zerolog-gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/stlimtat/sqxlab/go/internal/config"
	shttp "github.com/stlimtat/sqxlab/go/internal/http"
	"github.com/stlimtat/sqxlab/go/internal/telemetry"
	"golang.org/x/sync/errgroup"
)

type serverCmd struct {
	cmd    *cobra.Command
	server *Server
}

func newServerCmd(ctx context.Context) (*serverCmd, *cobra.Command) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("Testing")
	var err error

	serverCmd := &serverCmd{}

	// serverCmd represents the server command
	serverCmd.cmd = &cobra.Command{
		Use:   "server",
		Short: "Run the sqxlab cdp server",
		Long:  `Runs the Chrome Devtools Protocol server`,
		Args: func(_ *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCmd.server = newServer(cmd, args)
			err = serverCmd.server.Run(cmd.Context())
			return err
		},
	}

	return serverCmd, serverCmd.cmd
}

type Server struct {
	AdminSvr *http.Server
	Cfg      config.ServerConfig
	Gin      *gin.Engine
}

func newServer(
	cmd *cobra.Command,
	_ []string,
) *Server {
	ctx := cmd.Context()
	logger := zerolog.Ctx(ctx)
	var err error
	result := &Server{}

	result.Cfg = config.NewServerConfig(ctx)

	if result.Cfg.Debug {
		telemetry.SetGlobalLogLevel(zerolog.DebugLevel)
		gin.SetMode(gin.DebugMode)
	}
	result.Gin = gin.New()
	result.Gin.Use(gin.Recovery())
	result.Gin.Use(
		zerologgin.LoggerWithOptions(
			&zerologgin.Options{
				Name:   "sqxlab",
				Logger: logger,
			},
		),
	)
	err = shttp.RegisterAdminRoutes(ctx, result.Gin)
	if err != nil {
		logger.Fatal().Err(err).Msg("http.NewAdminRoutes")
	}

	result.AdminSvr = &http.Server{
		Addr:              ":8000",
		Handler:           result.Gin,
		ReadHeaderTimeout: 10 * time.Second,
	}

	return result
}

func (s *Server) Run(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	var err error

	eg.Go(func() error {
		err = s.AdminSvr.ListenAndServe()
		if err != nil {
			logger.Error().Err(err).Msg("AdminSvr.ListenAndServe")
		}
		return err
	})

	eg.Go(func() error {
		<-ctx.Done()
		ctx1, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		logger.Warn().Msg("Shutting down")
		err = s.AdminSvr.Shutdown(ctx1)
		if err != nil {
			logger.Error().Err(err).Msg("AdminSvr.Shutdown")
		}
		return err
	})

	err = eg.Wait()
	if err != nil {
		logger.Error().Err(err).Msg("errgroup Wait")
	}
	return err
}
