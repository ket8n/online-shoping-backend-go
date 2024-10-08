package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "online-shopping/models"
    "strconv"
)

func GetOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindUri(&order); err != nil {
        AsciiJSONError(c, "Invalid order ID")
        return
    }
    c.HTML(http.StatusOK, "order_summary.tmpl", gin.H{
        "order_id": order.ID,
    })
}

func FetchItems(c *gin.Context) {
    category := c.Query("category")
    price, _ := strconv.Atoi(c.Query("price"))

    c.JSON(http.StatusOK, gin.H{
        "category": category,
        "price":    price,
    })
}
