package config

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type ScreencastConfig struct {
	Debug        bool               `json:"debug" mapstructure:"debug"`
	Session      SessionConfig      `json:"session" mapstructure:"session"`
	UrlDiscovery UrlDiscoveryConfig `json:"urls" mapstructure:"urls"`
}

func NewScreencastConfig(ctx context.Context) ScreencastConfig {
	logger := zerolog.Ctx(ctx)

	var result ScreencastConfig
	err := viper.Unmarshal(&result)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unmarshal")
	}

	logger.Info().
		Interface("viper.AllSettings", viper.AllSettings()).
		Interface("result", result).
		Msg("ScreencastConfig init")

	return result
}
