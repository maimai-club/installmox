package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/lpxelinux.0", "./lpxelinux.0")
	r.StaticFile("/linux26", "./pxeboot/linux26")
	r.StaticFile("/initrd", "./pxeboot/initrd")
	r.POST("/answer",Answer)
	r.Run()
}