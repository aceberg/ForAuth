package web

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
	"github.com/aceberg/ForAuth/internal/notify"
)

func loginHandler(c *gin.Context) {
	var target, name string

	proxyAddr := c.MustGet("proxyAddr").(string)
	targetStruct, ok := targetMap[proxyAddr]

	if ok {
		target = targetStruct.Target
		name = targetStruct.Name
	} else {
		target = appConfig.Target
		name = "Default"
	}

	authOk := auth.Auth(c, &authConf)
	if authOk {
		reverseProxy(c, target)
	} else {
		loginScreen(c, target, name)
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

func loginScreen(c *gin.Context, target, name string) {
	var guiData models.GuiData

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {

		msg := "User '" + username + "' logged in from " + c.Request.RemoteAddr + ". Session expires in " + authConf.Expire.String() + ". Target: " + target + " (" + name + ")"
		log.Println("INFO:", msg)
		go notify.Shout("ForAuth: "+msg, appConfig.Notify)

		auth.StartSession(c)

		c.Redirect(http.StatusFound, "/")
	} else {
		if username != "" {
			msg := "Incorrect login attempt by '" + username + "' with password '" + password + "' logged in from " + c.Request.RemoteAddr + ". Target: " + target + " (" + name + ")"
			log.Println("WARNING:", msg)
			go notify.Shout("ForAuth: "+msg, appConfig.Notify)
		}

		guiData.Config = appConfig

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "login.html", guiData)
	}
}
