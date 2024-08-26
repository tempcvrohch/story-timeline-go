package config

import (
	"eksrow.com/story-timeline-go/global"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() *global.Config {
	v := viper.New()
	v.SetConfigName("env")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetConfigType("yaml")

	var config global.Config

	if err := v.ReadInConfig(); err != nil {
		log.Error().Msgf("viper:ReadInConfig: %v", err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Msg("viper:OnConfigChange")
	})

	if err := v.Unmarshal(&config); err != nil {
		log.Error().Msgf("viper:Unmarshal: %v", err)
	}

	// Gemini.Model not loading: https://github.com/spf13/viper/issues/769
	defaultConfigs := viper.AllSettings()
	for k, v := range defaultConfigs {
		viper.SetDefault(k, v)
	}
	v.WatchConfig()

	return &config
}
