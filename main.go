package main

import (
    "github.com/gin-gonic/gin"
    "online-shopping/routes"
    "online-shopping/middlewares"
)

func main() {
    r := gin.Default()

    r.Use(middlewares.RequestLogger())

    routes.InitRoutes(r)

	r.Run(":5000")
}
