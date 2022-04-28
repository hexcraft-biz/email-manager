package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/email-manager/authu"
	"github.com/hexcraft-biz/email-manager/config"
	"github.com/hexcraft-biz/email-manager/content"
	"github.com/hexcraft-biz/email-manager/email"
	"net/http"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err.Error())
	}

	_, err = conf.GetDB()
	defer conf.CloseDB()
	if err != nil {
		panic(err.Error())
	}

	SetRouter(conf).Run(":" + conf.AppPort)
}

func SetRouter(conf *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(MiddlewareFunc(conf))

	r.GET("/ping", Ping())
	r.GET("/email", email.Get(conf))
	r.POST("/email", email.Add(conf))
	r.PATCH("/email/count", authu.GoogleCronValidation(conf), email.ClearCounter(conf))
	r.POST("/content", content.Send(conf))

	return r
}

//================================================================
//
//================================================================
func MiddlewareFunc(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

//================================================================
//
//================================================================
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ggininder"})
	}
}
