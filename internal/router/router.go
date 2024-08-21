package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api"
	"github.com/hespecial/gin-mall/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	// http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Static("/storage", "./storage")

	v1 := router.Group("/api/v1")
	{
		// 测试接口
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})

		// 认证
		auth := v1.Group("/auth")
		{
			auth.POST("/register", api.Register)
			auth.POST("/login", api.Login)
		}

		// 用户操作
		authed := v1.Group("")
		authed.Use(middleware.JWTAuthMiddleware())
		{
			user := authed.Group("/user")
			{
				user.GET("/info", api.ShowUserInfo)
				user.PUT("/info", api.UserInfoUpdate)
				user.PUT("/password", api.UserPasswordChange)
				user.POST("/avatar", api.UploadAvatar)
			}
		}
	}

	return router
}

func Run() {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port),
		Handler: newRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listen failed: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server listen failed: %v", err)
	}
}
