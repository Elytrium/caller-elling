package types

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

var ModuleConfig Config

type Config struct {
	RequestExpiration int64 `default:"120"`
}

func LoadConfig() {
	err := envconfig.Process("caller", &ModuleConfig)
	if err != nil {
		log.Error().Err(err).Send()
	}
}
