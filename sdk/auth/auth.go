package auth

import "github.com/gin-gonic/gin"

func GetIDFromBearer(c *gin.Context) int {
	return c.MustGet("login").(int)
}
