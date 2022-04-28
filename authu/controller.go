package authu

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/email-manager/config"
	"net/http"
)

//================================================================
//
//================================================================
func GoogleCronValidation(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Xgc-Authorization") != conf.XgcAuthorization {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		}
	}
}
