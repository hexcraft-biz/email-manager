package content

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hexcraft-biz/email-manager/config"
	"github.com/hexcraft-biz/email-manager/model"
	"net/http"
)

type ReqSend struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

//================================================================
//
//================================================================
func Send(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := model.NewEEmail(conf.DB)
		if email, err := e.GetEmailWithLowestCount(); err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": model.ErrEmailReachedQuotaLimit.Error()})
		} else {
			rs := new(ReqSend)
			if err := c.ShouldBindWith(rs, binding.JSON); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			} else if len(rs.To) > model.DefGmailReceiverPerMailLimit {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "To must <= 5"})
			} else {
				if err := email.Send(conf, rs.To, rs.Subject, rs.Body); err != nil {
					c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": err.Error()})
				} else {
					if err := e.HitDailyCount(email.IDBin); err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
					} else {
						c.JSON(http.StatusCreated, gin.H{"message": http.StatusText(http.StatusCreated)})
					}
				}
			}
		}
	}
}
