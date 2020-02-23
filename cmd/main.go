package main

import (
	"github.com/integram-org/gitlab"
	"github.com/kelseyhightower/envconfig"
	"github.com/requilence/integram"
)

func main() {
	var cfg gitlab.Config
	envconfig.MustProcess("GITLAB", &cfg)

	integram.Register(
		cfg,
		cfg.BotConfig.Token,
	)

	integram.Run()
}
