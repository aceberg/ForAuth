package web

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
)

func loginHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		reverseProxy(c)
	} else {
		loginScreen(c)
	}
}

func reverseProxy(c *gin.Context) {

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = appConfig.Target
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func loginScreen(c *gin.Context) {
	var guiData models.GuiData

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {

		log.Println("INFO: user '"+username+"' logged in. Session expire time", authConf.Expire)

		auth.StartSession(c)

		c.Redirect(http.StatusFound, "/")
	} else {
		guiData.Config = appConfig

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "login.html", guiData)
	}
}
