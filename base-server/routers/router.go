package routers

import (
	"base-server/middleware/jwt"
	"base-server/routers/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

//-----------------------------------------------------------------------------

// InitRouter 初使化router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/wav", http.Dir("./audio"))

	r.POST("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/labelTotal", api.GetTotalLabel)
		apiV1.GET("/label/:index", api.GetLabelByIdx)
		apiV1.GET("/labels", api.ListLabel)
		apiV1.POST("/label", api.AddLabel)
	}

	return r
}
