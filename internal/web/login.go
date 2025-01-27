package web

import (
	"log"
	// "fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
	"github.com/aceberg/ForAuth/internal/notify"
)

func loginHandler(c *gin.Context) {
	var target string
	var ok bool

	reqHost := c.Request.Host
	target, ok = targetMap[reqHost]
	if !ok {
		target = appConfig.Target
	}

	authOk := auth.Auth(c, &authConf)
	if authOk {
		reverseProxy(c, target)
	} else {
		loginScreen(c, target)
	}
}

func reverseProxy(c *gin.Context, target string) {

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func loginScreen(c *gin.Context, target string) {
	var guiData models.GuiData

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {

		msg := "User '" + username + "' logged in. Session expires in " + authConf.Expire.String() + ". Target: " + target
		log.Println("INFO:", msg)
		notify.Shout("ForAuth: "+msg, appConfig.Notify)

		auth.StartSession(c)

		c.Redirect(http.StatusFound, "/")
	} else {
		guiData.Config = appConfig

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "login.html", guiData)
	}
}
