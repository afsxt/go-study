package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("中间件开始执行了")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		log.Println("中间件执行完毕", status, " time: ", time.Since(t))
	}
}

//-----------------------------------------------------------------------------
func main() {
	r := gin.Default()
	r.Use(MiddleWareTest())

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/welcome", func(c *gin.Context) {
		hStr := c.DefaultQuery("hello", "hello")
		req, _ := c.Get("request")
		log.Println("request: ", req)
		c.String(http.StatusOK, hStr)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			//c.ShouldBindJSON()
			c.JSON(http.StatusOK, gin.H{"success": "test ok"})
		})
	}

	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步执行: " + copyContext.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步执行: " + c.Request.URL.Path)
	})

	r.Run(":8000")
}
