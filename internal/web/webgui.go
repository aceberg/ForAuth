package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/check"
	"github.com/aceberg/ForAuth/internal/conf"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig, authConf = conf.Get(confPath)
	appConfig.DirPath = dirPath
	appConfig.ConfPath = confPath
	if nodePath != "" {
		appConfig.NodePath = nodePath
	}
	appConfig.Icon = icon

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	addressProxy := appConfig.Host + ":" + appConfig.Port
	addressConf := appConfig.Host + ":" + appConfig.PortConf

	log.Println("=================================== ")
	log.Printf("Config at http://%s", addressConf)
	log.Println("=================================== ")
	log.Printf("Proxy at http://%s", addressProxy)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	routerProxy := gin.Default()
	routerConf := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	routerProxy.SetHTMLTemplate(templ)          // templates
	routerConf.SetHTMLTemplate(templ)           // templates
	routerConf.StaticFS("/fs/", http.FS(pubFS)) // public

	routerProxy.GET("/*any", loginHandler)  // login.go
	routerProxy.POST("/*any", loginHandler) // login.go

	routerConf.GET("/", configHandler)              // config.go
	routerConf.GET("/logout", logoutHandler)        // config.go
	routerConf.POST("/", configHandler)             // config.go
	routerConf.POST("/config/", saveConfigHandler)  // config.go
	routerConf.POST("/config/auth", saveConfigAuth) // config.go

	go func() {
		err := routerConf.Run(addressConf)
		check.IfError(err)
	}()

	err := routerProxy.Run(addressProxy)
	check.IfError(err)
}
