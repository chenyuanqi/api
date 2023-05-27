package middlewares

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func CrossDomain(c *gin.Context) {
	//origin
	allowOrigin := c.Request.Header.Get("Origin")
	if allowOrigin == "" {
		if referer := c.Request.Header.Get("Referer"); referer != "" {
			u, _ := url.ParseRequestURI(referer)
			allowOrigin = u.Host
		}
	}

	c.Header("Access-Control-Allow-Origin", allowOrigin)
	c.Header("Access-Control-Allow-Headers", "DNT, X-Mx-ReqToken, Keep-Alive, X-Requested-With, Cache-Control, If-Modified-Since, token,access-token, X-Origin, Origin, Accept, Content-Type, Referer, User-Agent, Cookie, access-token, crossdomain, withCredentials, authorization")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("x-container-pod", os.Getenv("HOSTNAME"))

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
