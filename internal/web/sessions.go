package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
)

func delSessionHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		key := c.Query("key")
		auth.LogOutByToken(key)
		c.Redirect(http.StatusFound, "/sessions")
	}
}

func sessionsHandler(c *gin.Context) {

	authOk := auth.Auth(c, &authConf)
	if authOk {
		var guiData models.GuiData

		guiData.Config = appConfig
		guiData.Sessions = auth.GetAllSessions()

		c.HTML(http.StatusOK, "header.html", guiData)
		c.HTML(http.StatusOK, "sessions.html", guiData)
	} else {
		var targetStruct models.TargetStruct
		targetStruct.Target = appConfig.Host + ":" + appConfig.PortConf
		targetStruct.Name = "Config"

		loginScreen(c, targetStruct) // login.go
	}
}
