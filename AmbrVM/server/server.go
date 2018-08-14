package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() error {

	router := gin.Default()
	//router.Static("/views", "./server")
	//router.StaticFS("/views", http.Dir("server/views"))
	//http.Handle("/memfs/", http.StripPrefix("/memfs/", http.FileServer(fs)))

	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		buf, err := ioutil.ReadFile("server/views/index.html")
		if err != nil {
			c.Writer.WriteString(err.Error())
			return
		}
		c.Writer.Write(buf)
		//c.Redirect(http.StatusMovedPermanently, "/views/index.html")
	})

	router.GET("/index_cn.html", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "views/index_CN.html")
	})

	router.Run(":80")
	return nil
}
