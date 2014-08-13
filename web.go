package main

import (
    "io/ioutil"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", hello)
    r.GET("/ping", ping)

    r.Run(":8080")
}

func hello(c *gin.Context) {
    content, err := ioutil.ReadFile("main.html")
    if err != nil {
        //Do something
    }
    c.Data(200, "text/html", content)
}

func ping(c *gin.Context) {
    c.String(200, "Pong")
}
