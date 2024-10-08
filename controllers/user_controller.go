package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "online-shopping/models"
    "online-shopping/utils"
    "github.com/go-playground/validator/v10"
    "github.com/gin-gonic/gin/binding"
    // "encoding/json"
    // "github.com/gin-contrib/sse"
)

func AsciiJSONError(c *gin.Context, message string) {
    c.AsciiJSON(http.StatusBadRequest, gin.H{"error": message})
}

func RegisterUser(c *gin.Context) {
    var user models.User
    v := validator.New()
    v.RegisterValidation("password", utils.ValidatePassword)

    if err := c.ShouldBindWith(&user, binding.FormPost); err != nil {
        AsciiJSONError(c, "Invalid form input")
        return
    }

    if err := v.Struct(user); err != nil {
        AsciiJSONError(c, "Invalid email or password")
        return
    }

    c.AsciiJSON(http.StatusOK, gin.H{"status": "user registered"})
}
