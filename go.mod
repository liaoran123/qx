module eq

go 1.20

require xbdb v0.0.0

replace xbdb => ../xbdb

//go mod tidy
require (
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
)
