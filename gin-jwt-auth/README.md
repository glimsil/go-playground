# gin jwt auth server
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

You need to install the following modules:
    go get -u github.com/gin-gonic/gin
    go get -u github.com/gin-gonic/contrib/jwt
    go get -u github.com/dgrijalva/jwt-go
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/gin-swagger/swaggerFiles

To generate swagger files under the docs folder you need to type (you can check swaggo repository https://github.com/swaggo/swag):
    swag init

This services uses mongo as database, so you need a running mongo on port 27017. You can runnit with docker, using the following command:
    docker run -d --name some-mongo -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -p 27017:27017 mongo

Configure a database and add grants to a specific user. Then, add the connection informations to the config file(config/config.json)

Then, to execute the application, you can just type:
    go run main.go

Service will listen on port 8080.

You can access swagger on http://localhost:8080/swagger/index.html