package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitDBHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before request handler called")
		c.Next()
		fmt.Println("after request handler called")
	}
}
