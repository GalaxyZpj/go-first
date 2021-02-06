package main

import (
	g "github.com/gin-gonic/gin"
	. "github.com/galaxyzpj/gofirst/src/api"
)


func main() {
	router := g.Default()

	router.GET("", TestGETEndpoint)
	router.POST("", TestPOSTEndpoint)
	router.PUT("", TestPUTEndpoint)
	router.PATCH("", TestPATCHEndpoint)
	router.DELETE("", TestDELETEEndpoint)

	router.NoRoute(NoRoute)
	
	g.SetMode(g.ReleaseMode)

	router.Run("localhost:8000")
}
