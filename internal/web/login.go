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
	var authOk bool

	proxyAddr := c.MustGet("proxyAddr").(string)
	targetStruct, ok := targetMap[proxyAddr]

	if !ok {
		targetStruct.Target = appConfig.Target
		targetStruct.Name = "Default"

		authOk = auth.Auth(c, &authConf)
	} else {
		username, sesOk := auth.GetCurrentUser(c)
		targetUser, ok := targetStruct.Users[username]

		if sesOk && ((ok && targetUser.Auth) || username == authConf.User) {
			authOk = true
		}
	}

	if authOk {
		reverseProxy(c, targetStruct.Target)
	} else {
		loginScreen(c, targetStruct)
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

func loginScreen(c *gin.Context, targetStruct models.TargetStruct) {
	var guiData models.GuiData
	var client models.ClienInfo
	var clientStr string

	username := c.PostForm("username")
	password := c.PostForm("password")

	client.IP = c.PostForm("client_ip")
	client.Hostname = c.PostForm("client_hn")
	client.Timezone = c.PostForm("client_tz")
	client.City = c.PostForm("client_ci")
	client.Country = c.PostForm("client_co")

	if client.IP != "" {
		clientStr = "\nUser IP Info: \nIP: " + client.IP +
			"\nHostname: " + client.Hostname + "\nTimezone: " + client.Timezone + "\nCity: " + client.City + "\nCountry: " + client.Country
	}

	currentAuth, ok := checkUsername(targetStruct, username, password)

	if ok {

		msg := "User '" + username + "' logged in from " + c.Request.Host + ". Session expires in " + currentAuth.Expire.String() + ". Target: " + targetStruct.Target + " (" + targetStruct.Name + ")" + clientStr
		log.Println("INFO:", msg)
		go notify.Shout(msg, appConfig.Notify)
		go notify.Shout(msg, appConfig.Notify2)

		auth.StartSession(c, currentAuth, client.IP, targetStruct.Name)

		c.Redirect(http.StatusFound, c.Request.Referer())
	} else {
		if username != "" {
			msg := "Incorrect login attempt by '" + username + "' with password '" + password + "' logged in from " + c.Request.Host + ". Target: " + targetStruct.Target + " (" + targetStruct.Name + ")" + clientStr
			log.Println("WARNING:", msg)
			go notify.Shout(msg, appConfig.Notify)
			go notify.Shout(msg, appConfig.Notify2)
		}

		guiData.Config = appConfig

		c.HTML(http.StatusOK, "login.html", guiData)
	}
}

func checkUsername(targetStruct models.TargetStruct, username, password string) (auth.Conf, bool) {

	if username == authConf.User && auth.MatchPasswords(authConf.Password, password) {
		return authConf, true
	}

	targetAuth, ok := targetStruct.Users[username]
	if ok && (username == targetAuth.User && targetAuth.Auth &&
		auth.MatchPasswords(targetAuth.Password, password)) {

		targetAuth.Expire = auth.ToTime(targetAuth.ExpStr)
		return targetAuth, true
	}

	return authConf, false
}
