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
	routerConf := gin.New()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))

	routerConf.SetHTMLTemplate(templ)           // templates
	routerConf.StaticFS("/fs/", http.FS(pubFS)) // public

	routerConf.GET("/", configHandler)               // config.go
	routerConf.GET("/logout", logoutHandler)         // config.go
	routerConf.GET("/target/del", delTargetHandler)  // config.go
	routerConf.POST("/", configHandler)              // config.go
	routerConf.POST("/config/", saveConfigHandler)   // config.go
	routerConf.POST("/config/auth", saveConfigAuth)  // config.go
	routerConf.POST("/target/add", addTargetHandler) // config.go

	if appConfig.Port != "" {
		proxy := appConfig.Host + ":" + appConfig.Port
		go newRouter(templ, proxy, appConfig.Target, "Default")
	}

	for proxy, target := range targetMap {
		go newRouter(templ, proxy, target.Target, target.Name)
	}

	err := routerConf.Run(addressConf)
	check.IfError(err)
}

func newRouter(templ *template.Template, proxy, target, name string) {

	routerProxy := gin.New()
	routerProxy.SetHTMLTemplate(templ) // templates

	// Middleware to add variable to context
	routerProxy.Use(func(c *gin.Context) {
		c.Set("proxyAddr", proxy)
		c.Next()
	})

	routerProxy.GET("/*any", loginHandler)     // login.go
	routerProxy.POST("/*any", loginHandler)    // login.go
	routerProxy.PUT("/*any", loginHandler)     // login.go
	routerProxy.DELETE("/*any", loginHandler)  // login.go
	routerProxy.PATCH("/*any", loginHandler)   // login.go
	routerProxy.HEAD("/*any", loginHandler)    // login.go
	routerProxy.OPTIONS("/*any", loginHandler) // login.go

	log.Println("Proxy at http://"+proxy, "=> http://"+target, "("+name+")")
	log.Println("=================================== ")

	err := routerProxy.Run(proxy)
	check.IfError(err)
}
