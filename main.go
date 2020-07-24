package main

import (
	"context"
	"fmt"
	"github.com/HLSO6/go-gin-example/pkg/setting"
	"github.com/HLSO6/go-gin-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//linux下使用endless
	////优雅的关闭服务
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HttpPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"message":"test",
	//	})
	//})
	//路由独立出来
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//s.ListenAndServe()
	//优雅关闭 http.Server 的 Shutdown 方法
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
