package utils

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
func getenv(key, fallBack string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallBack
	}

	return value
}

var SECRET_KEY = []byte(getenv("SECRET_KEY", "tranquanthuan132@gmail.com"))
var RUNNING_MODE = getenv("RUNNING_MODE", gin.DebugMode)
var PORT = getenv("PORT", "8080")
