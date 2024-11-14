package main

import (
	"fmt"
	"github.com/bianpeijiang/ginblog/pkg/setting"
	"github.com/bianpeijiang/ginblog/routers"
	"net/http"
)

func main() {
	// router := gin.Default()
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic("start server error:" + err.Error())
	}
}
