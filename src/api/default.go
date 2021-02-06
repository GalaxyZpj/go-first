package api

import (
	. "github.com/gin-gonic/gin"
	HTTP "net/http"
)

func NoRoute(c *Context) {
	c.JSON(HTTP.StatusNotFound, H{
		"status": "Not Found",
		"message": "Specified route not found.",
	})
}
