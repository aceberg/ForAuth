package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ForAuth/internal/check"
	"github.com/aceberg/ForAuth/internal/conf"
	"github.com/aceberg/ForAuth/internal/yaml"
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

	appConfig.YamlPath = dirPath + "/targets.yaml"
	check.Path(appConfig.YamlPath)
	targetMap = yaml.Read(appConfig.YamlPath)

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	addressConf := appConfig.Host + ":" + appConfig.PortConf

	log.Println("=================================== ")
	log.Println("Config at http://" + addressConf)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	routerProxy := gin.Default()
	routerConf := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	routerProxy.SetHTMLTemplate(templ)          // templates
	routerConf.SetHTMLTemplate(templ)           // templates
	routerConf.StaticFS("/fs/", http.FS(pubFS)) // public

	routerProxy.GET("/*any", loginHandler)     // login.go
	routerProxy.POST("/*any", loginHandler)    // login.go
	routerProxy.PUT("/*any", loginHandler)     // login.go
	routerProxy.DELETE("/*any", loginHandler)  // login.go
	routerProxy.PATCH("/*any", loginHandler)   // login.go
	routerProxy.HEAD("/*any", loginHandler)    // login.go
	routerProxy.OPTIONS("/*any", loginHandler) // login.go

	routerConf.GET("/", configHandler)               // config.go
	routerConf.GET("/logout", logoutHandler)         // config.go
	routerConf.GET("/target/del", delTargetHandler)  // config.go
	routerConf.POST("/", configHandler)              // config.go
	routerConf.POST("/config/", saveConfigHandler)   // config.go
	routerConf.POST("/config/auth", saveConfigAuth)  // config.go
	routerConf.POST("/target/add", addTargetHandler) // config.go

	if appConfig.Port != "" {
		proxy := appConfig.Host + ":" + appConfig.Port
		log.Println("Proxy at http://"+proxy, "=> http://"+appConfig.Target)
		log.Println("=================================== ")
		go func() {
			err := routerProxy.Run(proxy)
			check.IfError(err)
		}()
	}

	for proxy, target := range targetMap {
		log.Println("Proxy at http://"+proxy, "=> http://"+target.Target)
		log.Println("=================================== ")
		go func() {
			err := routerProxy.Run(proxy)
			check.IfError(err)
		}()
	}

	err := routerConf.Run(addressConf)
	check.IfError(err)
}
