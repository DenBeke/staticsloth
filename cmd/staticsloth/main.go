package main

import (
	"github.com/DenBeke/staticsloth"
	log "github.com/sirupsen/logrus"
)

func main() {

	config := staticsloth.BuildConfigFromEnv()

	err := config.Validate()
	if err != nil {
		log.Fatalf("config file invalid: %v", err)
	}

	log.Printf("starting StaticSloth with config: %+v", config)

	staticsloth.Serve(config)

}
