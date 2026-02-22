package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
	"github.com/aceberg/ForAuth/internal/yaml"
)

func advancedHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		var guiData models.GuiData

		targetMap = yaml.Read(appConfig.YamlPath)

		guiData.Config = appConfig
		guiData.TargetMap = targetMap

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "advanced.html", guiData)
	} else {
		var targetStruct models.TargetStruct
		targetStruct.Target = appConfig.Host + ":" + appConfig.PortConf
		targetStruct.Name = "Config"

		loginScreen(c, targetStruct) // login.go
	}
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
	c.Redirect(http.StatusFound, c.Request.Referer())
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

func addUserHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)

	target := c.PostForm("target")
	user := c.PostForm("user")
	password := c.PostForm("password")
	expire := c.PostForm("expire")

	if authOk && user != "" && password != "" {

		if expire == "" {
			expire = "14d"
		}

		newUser := auth.Conf{Auth: true, User: user, ExpStr: expire}
		newUser.Password = auth.HashPassword(password)

		proxy := targetMap[target]

		userMap := make(map[string]auth.Conf)
		if proxy.Users != nil {
			userMap = proxy.Users
		}
		userMap[user] = newUser
		proxy.Users = userMap

		targetMap[target] = proxy

		yaml.Write(appConfig.YamlPath, targetMap)
	}
	c.Redirect(http.StatusFound, c.Request.Referer())
}

func delUserHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		target := c.Query("target")
		user := c.Query("user")

		proxy := targetMap[target]
		userMap := proxy.Users

		delete(userMap, user)
		proxy.Users = userMap
		targetMap[target] = proxy

		yaml.Write(appConfig.YamlPath, targetMap)
	}
	c.Redirect(http.StatusFound, c.Request.Referer())
}

func enableUserHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		target := c.Query("target")
		user := c.Query("user")

		proxy := targetMap[target]
		userMap := proxy.Users

		newUser := userMap[user]
		newUser.Auth = !newUser.Auth

		userMap[user] = newUser
		proxy.Users = userMap
		targetMap[target] = proxy

		yaml.Write(appConfig.YamlPath, targetMap)
	}
	c.Redirect(http.StatusFound, c.Request.Referer())
}
