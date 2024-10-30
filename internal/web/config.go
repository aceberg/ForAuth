package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/conf"
	"github.com/aceberg/ForAuth/internal/models"
)

func configHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig
	guiData.Auth = authConf

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "config.html", guiData)
}

func saveConfigHandler(c *gin.Context) {

	appConfig.Host = c.PostForm("host")
	appConfig.Port = c.PostForm("port")
	appConfig.Theme = c.PostForm("theme")
	appConfig.Color = c.PostForm("color")

	conf.Write(appConfig, authConf)

	log.Println("INFO: writing new config to", appConfig.ConfPath)

	c.Redirect(http.StatusFound, "/")
}

func saveConfigAuth(c *gin.Context) {

	authConf.User = c.PostForm("user")
	authConf.ExpStr = c.PostForm("expire")
	authStr := c.PostForm("auth")
	pw := c.PostForm("password")

	if authStr == "on" {
		authConf.Auth = true
	} else {
		authConf.Auth = false
	}
	appConfig.Auth = authConf.Auth

	if pw != "" {
		authConf.Password = auth.HashPassword(pw)
	}

	authConf.Expire = auth.ToTime(authConf.ExpStr)

	if authConf.Auth && (authConf.User == "" || authConf.Password == "") {
		log.Println("WARNING: Auth won't work with empty login or password.")
		authConf.Auth = false
	}

	conf.Write(appConfig, authConf)

	c.Redirect(http.StatusFound, "/")
}
