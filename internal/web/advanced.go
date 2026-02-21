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

		guiData.Config = appConfig
		guiData.TargetMap = yaml.Read(appConfig.YamlPath)

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
