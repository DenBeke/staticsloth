# 🦥 StaticSloth

<!-- [![Build Status](https://travis-ci.com/DenBeke/staticsloth.svg?branch=master)](https://travis-ci.com/DenBeke/staticsloth) -->
[![Go Report Card](https://goreportcard.com/badge/github.com/DenBeke/staticsloth)](https://goreportcard.com/report/github.com/DenBeke/staticsloth)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/denbeke/staticsloth?sort=date)](https://hub.docker.com/r/denbeke/staticsloth)

⚠️ **WIP**

A static file server written in Go.
All configuration is done with environment variables.


## Usage

    go run cmd/staticsloth/*.go 


## Variables

`HTTP_ADDRESS`: HTTP listen address. Default: `:1234`

`PATH_PREFIX`: Path prefix that StaticSloth will serve for. Default: `/`

`DIRECTORY`: Directory that will be served. Default: `/var/www/html`

`ACCESS_LOG`: Displays access logs when set to `1`. Default: `0`

`GZIP`: Compress responses with gzip. Default: `1`

`CACHE_CONTROL_PATHS`: List of (comma separated) paths to set `Cache-Control: max-age=X` HTTP header for. Example: `CACHE_CONTROL_PATHS="/assets,/images,/uploads"` Default: none (= disabled)

`CACHE_CONTROL_DURATION`: Duration value (in seconds) that will be set in the `Cache-Control` header. Default: none (= disabled)

`BLOCK_PATHS`: List of (comma separated) paths to block. Example: `BLOCK_PATHS="/.git"` Default: (= `/.git`)


## Multi architecture Docker build

    docker buildx build --platform linux/amd64,linux/arm64,linux/arm --tag denbeke/staticsloth --push


## Author

[Mathias Beke](https://denbeke.be)