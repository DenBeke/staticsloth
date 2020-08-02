package main

import (
	"github.com/DenBeke/staticsloth"
	log "github.com/sirupsen/logrus"
)

func main() {

	config := staticsloth.BuildConfigFromEnv()

	err := config.Validate()
	if err != nil {
		log.Fatalln("config file invalid: %v", err)
	}

	staticsloth.Serve(config)

}
