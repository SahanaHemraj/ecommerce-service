package user

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func SignUp(c *gin.Context) {
    // Logic for signing up (add user to DB)
}

func Login(c *gin.Context) {
    // Logic for login (validate user and return JWT token)
}
