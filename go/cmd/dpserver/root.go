/*
Copyright © 2024 Swee Tat Lim <st_lim@stlim.net>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stlimtat/sqxlab/go/internal/config"
	"github.com/stlimtat/sqxlab/go/internal/telemetry"
)

type rootCmd struct {
	cmd *cobra.Command
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func newRootCmd(ctx context.Context) *rootCmd {
	logger := zerolog.Ctx(ctx)
	var err error

	result := &rootCmd{}
	// rootCmd represents the base command when called without any subcommands
	result.cmd = &cobra.Command{
		Use:   "dpserver",
		Short: "An Chrome DevTools Protocol server",
		Long:  "An Chrome DevTools Protocol server",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			cmdCtx, _ := telemetry.GetLogger(cmd.Context(), cmd.OutOrStdout())
			cmd.SetContext(cmdCtx)
		},
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(config.RootConfigInit)

	result.cmd.PersistentFlags().BoolP(
		"debug", "d",
		false,
		"Run the application in debug mode",
	)
	result.cmd.PersistentFlags().StringP(
		"config", "c",
		"",
		"config file (default is $HOME/config.yaml)",
	)
	err = viper.BindPFlag("debug", result.cmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		logger.Fatal().Err(err).Msg("viper.BindPFlag.debug")
	}
	err = viper.BindPFlag("config", result.cmd.PersistentFlags().Lookup("config"))
	if err != nil {
		logger.Fatal().Err(err).Msg("viper.BindPFlag.config")
	}
	_, serverCmd := newServerCmd(ctx)

	result.cmd.AddCommand(
		serverCmd,
	)

	return result
}

func (r *rootCmd) ExecuteContext(ctx context.Context) error {
	return r.cmd.ExecuteContext(ctx)
}
