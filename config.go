package staticsloth

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	defaultPort                 = 1234
	defaultHTTPAddress          = fmt.Sprintf(":%d", defaultPort)
	defaultPathPrefix           = "/"
	defaultDirectory            = "/var/www/html"
	defaultAccessLog            = false
	defaultGzip                 = true
	defaultCacheControlPaths    = []string{}
	defaultCacheControlDuration = -1
	defaultBlockPaths           = []string{"/.git"}
)

// BuildConfigFromEnv populates a StaticSloth config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.HTTPAddress = getEnv("HTTP_ADDRESS", defaultHTTPAddress)
	config.PathPrefix = getEnv("PATH_PREFIX", defaultPathPrefix)
	config.Directory = getEnv("DIRECTORY", defaultDirectory)

	// access log
	accessLog := getEnv("ACCESS_LOG", "0")
	if accessLog == "1" {
		config.AccessLog = true
	}

	// gzip
	gzip := getEnv("GZIP", "1")
	if gzip == "0" {
		config.Gzip = false
	} else {
		config.Gzip = true
	}

	// cache-control
	cacheControlPaths := getEnv("CACHE_CONTROL_PATHS", "")
	if cacheControlPaths == "" {
		config.CacheControlPaths = []string{}
	} else {
		config.CacheControlPaths = strings.Split(cacheControlPaths, ",")
	}

	cacheControlDuration := getEnv("CACHE_CONTROL_DURATION", "")
	if cacheControlDuration == "" {
		config.CacheControlDuration = defaultCacheControlDuration
	} else {
		duration, err := strconv.Atoi(cacheControlDuration)
		if err == nil {
			config.CacheControlDuration = duration
		}
	}

	// block paths
	blockPaths := getEnv("BLOCK_PATHS", "")
	if blockPaths == "" {
		config.BlockPaths = defaultBlockPaths
	} else {
		config.BlockPaths = strings.Split(blockPaths, ",")
	}

	return config
}

// Config contains all the config for serving a static site
type Config struct {
	HTTPAddress          string
	PathPrefix           string
	Directory            string
	AccessLog            bool
	Gzip                 bool
	CacheControlPaths    []string
	CacheControlDuration int
	BlockPaths           []string
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

	if len(config.CacheControlPaths) > 0 && config.CacheControlDuration == defaultCacheControlDuration {
		return fmt.Errorf("CacheControlDuration must be set if CacheControlPaths is set")
	}

	if len(config.CacheControlPaths) > 0 {
		for _, path := range config.CacheControlPaths {
			if !strings.HasPrefix(path, "/") {
				return fmt.Errorf("CacheControlPaths (%q) should start with '/'", path)
			}
		}
	}

	if len(config.BlockPaths) > 0 {
		for _, path := range config.BlockPaths {
			if !strings.HasPrefix(path, "/") {
				return fmt.Errorf("BlockPaths (%q) should start with '/'", path)
			}
		}
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
