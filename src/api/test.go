package api

import (
	. "github.com/gin-gonic/gin"
	HTTP "net/http"
)

func TestGETEndpoint(c *Context) {
	c.JSON(HTTP.StatusOK, H{
		"status": "OK",
		"message": "Hitting a test GET endpoint.",
	})
}

func TestPOSTEndpoint(c *Context) {
	c.JSON(HTTP.StatusCreated, H{
		"status": "OK",
		"message": "Hitting a test POST endpoint.",
	})
}

func TestPUTEndpoint(c *Context) {
	c.JSON(HTTP.StatusCreated, H{
		"status": "OK",
		"message": "Hitting a test PUT endpoint.",
	})
}

func TestPATCHEndpoint(c *Context) {
	c.JSON(HTTP.StatusOK, H{
		"status": "OK",
		"message": "Hitting a test PATCH endpoint.",
	})
}

func TestDELETEEndpoint(c *Context) {
	c.JSON(HTTP.StatusOK, H{
		"status": "OK",
		"message": "Hitting a test DELETE endpoint.",
	})
}
