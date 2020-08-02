# StaticSloth

⚠️ **WIP**

A static file server written in Go.
All configuration is done with environment variabels


## Usage

go run cmd/staticsloth/*.go 


## Variables

`HTTP_ADDRESS`: HTTP listen address. Default: `:1234`

`PATH_PREFIX`: Path prefix that StaticSloth will serve for. Default: `/`

`DIRECTORY`: Directory that will be served. Defaults: `/var/www/html`

`ACCESS_LOG`: Displays access logs when set to `1`. Defaults: `0`


## Author

[Mathias Beke](https://denbeke.be)