module github.com/HLSO6/go-gin-example

go 1.13

require (
	github.com/astaxie/beego v1.12.2
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.4 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
	golang.org/x/tools v0.0.0-20200721223218-6123e77877b2 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/HLSO6/go-gin-example/conf => ./go-gin-example/pkg/conf
	github.com/HLSO6/go-gin-example/middleware => ./go-gin-example/middleware
	github.com/HLSO6/go-gin-example/middleware/jwt => ./go-gin-example/middleware/jwt
	github.com/HLSO6/go-gin-example/models => ./go-gin-example/models
	github.com/HLSO6/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/HLSO6/go-gin-example/pkg/logging => ./go-gin-example/pkg/logging
	github.com/HLSO6/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/HLSO6/go-gin-example/pkg/util => ./go-gin-example/pkg/util
	github.com/HLSO6/go-gin-example/routers => ./go-gin-example/routers
	github.com/HLSO6/go-gin-example/routers/api/v1/ => ./go-gin-example/routers/api/v1/

)
