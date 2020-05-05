package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

// GetUser return dummy user info
func GetUser(c *gin.Context) {
	user := &User{
		Name: "test taro",
		Age:  20,
	}
	c.HTML(http.StatusOK, "user.tmpl", gin.H{
		"name": user.Name,
		"age":  user.Age,
	})
}
