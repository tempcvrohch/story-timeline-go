package main

import (
	"eksrow.com/story-timeline-go/config"
	"eksrow.com/story-timeline-go/global"
	"eksrow.com/story-timeline-go/logger"
	"eksrow.com/story-timeline-go/routes"
	"github.com/rs/zerolog/log"
)

var ctx *global.Context

func main() {
	logger.InitLogger()
	config := config.New()
	router := routes.New(config)

	ctx = &global.Context{Config: config, Router: router}
	log.Debug().Msgf("debug: %v", config.Debug)
}
