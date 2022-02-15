package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func loadPingRouter(group *gin.RouterGroup) {
	group.GET("/ping", ping)
}

func loadIndexRouter(group *gin.RouterGroup) {
	group.GET("", index)
	group.GET("/index.html", index)
	group.GET("/index.htm", index)
	group.GET("/index.jsp", index)
	group.GET("/index.php", index)
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func index(c *gin.Context) {
	msg := "Hello! This is used for view ip of pod."
	var hostIp []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		hostIp = append(hostIp, "no ip found")
	} else {
		for i := 0; i < len(netInterfaces); i++ {
			if (netInterfaces[i].Flags & net.FlagUp) != 0 {
				addrs, _ := netInterfaces[i].Addrs()
				for _, address := range addrs {
					if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							hostIp = append(hostIp, fmt.Sprintf("%s:%s", netInterfaces[i].Name, ipnet.IP.String()))
						}
					}
				}
			}
		}
	}

	hostIpText := strings.Join(hostIp, "\r\n")
	c.String(200, "%s\n\r\n%s\r\n\n%s", msg, hostIpText, "version:1.0.0")
}
