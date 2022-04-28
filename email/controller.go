package email

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hexcraft-biz/email-manager/config"
	"github.com/hexcraft-biz/email-manager/model"
	"net/http"
)

type ReqAdd struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

//================================================================
//
//================================================================
func Get(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if emails, err := model.NewEEmail(conf.DB).GetAllEnabled(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		} else if len(emails) <= 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK), "results": emails})
		}
	}
}

//================================================================
//
//================================================================
func Add(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ra := new(ReqAdd)
		if err := c.ShouldBindWith(ra, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return
		} else if len(ra.Password) > model.DefLenPassword {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": model.ErrEmailPasswordLen.Error()})
			return
		}

		if email, err := model.NewEEmail(conf.DB).Insert(ra.Addr, ra.Password); err != nil {
			switch err {
			case model.ErrEmailExists:
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": model.ErrEmailExists.Error()})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
			}
		} else {
			c.JSON(http.StatusCreated, gin.H{"message": http.StatusText(http.StatusCreated), "results": gin.H{"id": email.ID}})
		}
	}
}

//================================================================
//
//================================================================
func ClearCounter(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := model.NewEEmail(conf.DB).ResetDailyCount(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
		}
	}
}
