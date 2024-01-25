package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Description ping the server.
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	fmt.Println("masuk sini ga sih")
	c.JSON(http.StatusOK, "pong")
}
