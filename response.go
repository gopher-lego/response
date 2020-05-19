package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Normal logic right
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, jsonify(data, 0, "Success"))
}

// Normal logic fail
func Failure(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, jsonify(data, -1, "Failure"))
}

// Reject access with header status
func Reject(c *gin.Context, status int, data interface{}) {
	c.JSON(status, jsonify(data, -1, "Something Wrong"))
}

func jsonify(data interface{}, code int, msg string) gin.H {
	return gin.H{
		"data": data,
		"code": code,
		"msg": msg,
	}
}