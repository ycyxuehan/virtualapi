module bing89.com/virtualapi

go 1.14

require bing89.com/virtualapi/libs v0.0.0

require (
	bing89.com/virtualapi/admin v0.0.0
	github.com/gin-gonic/gin v1.6.3
)

replace bing89.com/virtualapi/libs v0.0.0 => ./libs

replace bing89.com/virtualapi/admin v0.0.0 => ./admin
