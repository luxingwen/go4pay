package middleware

import (
	"bytes"
	"go4pay/pkg/config"
	log "go4pay/pkg/logger"
	"go4pay/pkg/utils"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("Request Body:", string(b))

	// Write the body content back to the request
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

	c.Set("body", string(b))
	// Continue to next middleware or handler
	c.Next()
}

func OpenFixAuth(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if config.GetConfig().IsSign {

		reqSign := c.GetHeader("X-OpenPix-Signature")

		localSign := utils.HmacSha1Signature(b, []byte(config.GetConfig().SecretKeyOnOpenpixPlatform))

		if reqSign != string(localSign) {

			log.Errorf("签名验证不通过,reqSign:%s , localSign:%s", reqSign, string(localSign))
			c.Abort()
			return
		}
	} else {
		log.Errorf("注意:%s", "没有开启签名验证，将会有风险")
	}

	// Write the body content back to the request
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

	c.Next()
}
