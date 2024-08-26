package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hespecial/gin-mall/docs"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/initialize"
	"github.com/hespecial/gin-mall/internal/router"
)

func init() {
	global.Config = initialize.LoadConfig()
	global.DB = initialize.InitMySQL()
	global.EsClient = initialize.InitEs()
	global.Log = initialize.InitLogger()

	initialize.CreateDirectories()
	initialize.InitJWT()
	initialize.InitOSS()
	initialize.InitEmail()
}

//	@title			gin-mall
//	@version		1.0
//	@description	gin-mall API Documentation
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	1478488313@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	AccessToken
//	@in							header
//	@name						Authorization

//	@securityDefinitions.apikey	RefreshToken
//	@in							header
//	@name						X-Refresh-Token

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	gin.SetMode(global.Config.Server.Level)
	router.Run()
}
