package staticsloth

import (
	"fmt"
	"os"
	"strings"
)

var (
	defaultPort        = 1234
	defaultHTTPAddress = fmt.Sprintf(":%d", defaultPort)
	defaultPathPrefix  = "/"
	defaultDirectory   = "/var/www/html"
	defaultAccessLog   = false
)

// BuildConfigFromEnv populates a StaticSloth config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.HTTPAddress = getEnv("HTTP_ADDRESS", defaultHTTPAddress)
	config.PathPrefix = getEnv("PATH_PREFIX", defaultPathPrefix)
	config.Directory = getEnv("DIRECTORY", defaultDirectory)

	accessLog := getEnv("ACCESS_LOG", "0")
	if accessLog == "1" {
		config.AccessLog = true
	}

	return config
}

// Config contains all the config for serving a static site
type Config struct {
	HTTPAddress string
	PathPrefix  string
	Directory   string
	AccessLog   bool
}

// Validate validates whether all config is set and valid
func (config *Config) Validate() error {

	if config.HTTPAddress == "" {
		return fmt.Errorf("HTTPAddress cannot be empty")
	}
	if config.PathPrefix == "" {
		return fmt.Errorf("PathPrefix cannot be empty")
	}
	if !strings.HasPrefix(config.PathPrefix, "/") {
		return fmt.Errorf("PathPrefix should start with '/'. Got: %q", config.PathPrefix)
	}
	if config.Directory == "" {
		return fmt.Errorf("Directory cannot be empty")
	}

	return nil
}

// getEnv gets the env variable with the given key if the key exists
// else it falls back to the fallback value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
