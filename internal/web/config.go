package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/check"
	"github.com/aceberg/ForAuth/internal/conf"
	"github.com/aceberg/ForAuth/internal/models"
	"github.com/aceberg/ForAuth/internal/notify"
)

func logoutHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		auth.LogOut(c)
		c.Redirect(http.StatusFound, "/")
	}
}

func configHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		var guiData models.GuiData

		guiData.Config = appConfig
		guiData.Auth = authConf

		guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

		file, err := pubFS.ReadFile("public/version")
		check.IfError(err)
		version := string(file)
		guiData.Version = version[8:]

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "config.html", guiData)
	} else {
		var targetStruct models.TargetStruct
		targetStruct.Target = appConfig.Host + ":" + appConfig.PortConf
		targetStruct.Name = "Config"

		loginScreen(c, targetStruct) // login.go
	}
}

func saveConfigHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		appConfig.Host = c.PostForm("host")
		appConfig.Port = c.PostForm("port")
		appConfig.PortConf = c.PostForm("portconf")
		appConfig.Target = c.PostForm("target")
		appConfig.Theme = c.PostForm("theme")
		appConfig.Color = c.PostForm("color")
		appConfig.NodePath = c.PostForm("nodepath")
		appConfig.Notify = c.PostForm("notify")

		ipInfo := c.PostForm("ipinfo")
		if ipInfo == "on" {
			appConfig.IPInfo = true
		} else {
			appConfig.IPInfo = false
		}

		conf.Write(appConfig, authConf)

		log.Println("INFO: writing new config to", appConfig.ConfPath)
	}
	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveConfigAuth(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		authConf.User = c.PostForm("user")
		authConf.ExpStr = c.PostForm("expire")
		authStr := c.PostForm("auth")
		pw := c.PostForm("password")

		if authStr == "on" {
			authConf.Auth = true
		} else {
			authConf.Auth = false
		}

		if pw != "" {
			authConf.Password = auth.HashPassword(pw)
		}

		authConf.Expire = auth.ToTime(authConf.ExpStr)

		if authConf.Auth && (authConf.User == "" || authConf.Password == "") {
			log.Println("WARNING: Auth won't work with empty login or password.")
			authConf.Auth = false
		}

		log.Println("INFO: writing new auth config to", appConfig.ConfPath)
		conf.Write(appConfig, authConf)
	}

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func notifyHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {

		go notify.Shout("ForAuth: Test Notification", appConfig.Notify)

		c.Redirect(http.StatusFound, c.Request.Referer())
	}
}
