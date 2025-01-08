package config

type UrlDiscoveryConfig struct {
	URLs []string `json:"urls" mapstructure:"urls"`
}
