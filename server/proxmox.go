package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
)


func Answer(c *gin.Context){


	bytes, err := ioutil.ReadFile("./example.toml")
    if err != nil {
        panic(err)
    }

	c.String(http.StatusOK,string(bytes))

}