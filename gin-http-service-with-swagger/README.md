# gin http service with swagger
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

You need to install the following modules:
    go get -u github.com/gin-gonic/gin
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/gin-swagger/swaggerFiles

To generate swagger files under the docs folder you need to type (you can check swaggo repository https://github.com/swaggo/swag):
    swag init

Then you can just type:
    go run main.go

Server will listen on port 8080 with an endpoint on path `/ping`

You can access swagger on http://localhost:8080/swagger/index.html