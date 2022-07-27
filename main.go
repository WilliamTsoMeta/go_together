package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/json_map",func (c *gin.Context)  {
		c.JSON(200,map[string]interface{}{
			"success":true,
			"msg":"hello json",
		})
	})

	r.GET("/json_H",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"success":true,
			"msg":"hello jsonH",
		})
	})

	r.GET("/json_struct",func (c *gin.Context)  {
		type TestStruct struct {
			Title string
			Desc string
		}
		a := TestStruct{
			Title:"title",
			Desc:"decs",
		}
		c.JSON(200,a)
	})

	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}