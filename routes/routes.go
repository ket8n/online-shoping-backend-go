package routes

import (
    "github.com/gin-gonic/gin"
    "online-shopping/controllers"
)

func InitRoutes(r *gin.Engine) {
    r.POST("/register", controllers.RegisterUser)

    r.GET("/orders/:id", controllers.GetOrder)
    r.GET("/items", controllers.FetchItems)
}
