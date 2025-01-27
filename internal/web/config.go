package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/check"
	"github.com/aceberg/ForAuth/internal/conf"
	"github.com/aceberg/ForAuth/internal/models"
	"github.com/aceberg/ForAuth/internal/yaml"
)

func logoutHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		auth.LogOut(c)
		c.Redirect(http.StatusFound, "/")
	}
}

func configHandler(c *gin.Context) {
	var guiData models.GuiData

	authOk := auth.Auth(c, &authConf)
	if authOk {
		guiData.Config = appConfig
		guiData.Auth = authConf
		guiData.TargetMap = yaml.Read(appConfig.YamlPath)

		guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

		file, err := pubFS.ReadFile("public/version")
		check.IfError(err)
		version := string(file)
		guiData.Version = version[8:]

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "config.html", guiData)
	} else {
		loginScreen(c, appConfig.Host+":"+appConfig.PortConf) // login.go
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

		conf.Write(appConfig, authConf)

		log.Println("INFO: writing new config to", appConfig.ConfPath)
	}
	c.Redirect(http.StatusFound, "/")
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

		conf.Write(appConfig, authConf)
	}

	c.Redirect(http.StatusFound, "/")
}

func addTargetHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		name := c.PostForm("name")
		proxy := c.PostForm("proxy")
		target := c.PostForm("target")

		targetMap[proxy] = models.TargetStruct{Name: name, Target: target}
		yaml.Write(appConfig.YamlPath, targetMap)
	}
	c.Redirect(http.StatusFound, "/")
}

func delTargetHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		key := c.Query("key")

		delete(targetMap, key)
		yaml.Write(appConfig.YamlPath, targetMap)
	}
	c.Redirect(http.StatusFound, c.Request.Referer())
}
