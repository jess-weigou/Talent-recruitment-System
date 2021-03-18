package main

import "github.com/gin-gonic/gin"

func (s Service)initRouter()  {
    r:=gin.Default()
    //注册登陆模块
    r.GET("/login",s.Login)
    //r.POST("/register")
    //selfInfo:=r.Group("/selfInfo")
    //{
    //    selfInfo.POST("/add")
    //    selfInfo.GET("/see")
    //}
    s.Router=r
    r.Run(":8080")
}