package auth

import "github.com/gin-gonic/gin"

func GetIDFromBearer(c *gin.Context) int {
	id := c.MustGet("login").(int)
	return id
}
