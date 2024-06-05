package helpers

import "github.com/gin-gonic/gin"

// Mendapatkan content type dari request headers yang dikirim client
func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
