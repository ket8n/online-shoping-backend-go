package models

type User struct {
    Email    string `form:"email" binding:"required,email"`
    Password string `form:"password" binding:"required,min=8"`
}
