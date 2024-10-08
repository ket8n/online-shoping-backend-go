package middlewares

import (
    "time"

    "github.com/fatih/color"
    "github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        method := c.Request.Method
        path := c.Request.URL.Path

        c.Next()

        duration := time.Since(start)
        switch method {
        case "GET":
            color.Green("[GET] %s | %v", path, duration)
        case "POST":
            color.Blue("[POST] %s | %v", path, duration)
        case "DELETE":
            color.Red("[DELETE] %s | %v", path, duration)
        default:
            color.White("[%s] %s | %v", method, path, duration)
        }
    }
}
