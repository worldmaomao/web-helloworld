package internal

import (
	"worldmaomao/web-hellword/internal/rest"
)

func Initial() {

	// 启动web server
	rest.NewServer().Start()
}
