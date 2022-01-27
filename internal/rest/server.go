package rest

import (
	"github.com/gin-gonic/gin"
	"worldmaomao/web-hellword/internal/rest/middlewares"
)

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (server *server) Start() {
	address := "0.0.0.0:8080"
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middlewares.NewCors([]string{"*"}))
	// 不需要登录
	noAuthApiRoute := r.Group("/")
	loadPingRouter(noAuthApiRoute)
	loadIndexRouter(noAuthApiRoute)

	r.Run(address)
}
