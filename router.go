package main

import "github.com/gin-gonic/gin"

func (s Service)initRouter()  {
    r:=gin.Default()
    //注册登陆模块
    r.GET("/login",s.Login)
    r.POST("/register",s.Register)
    selfInfo:=r.Group("/profiles/:phone")
    {
       selfInfo.GET("/",s.GetSelfDetail)
       selfInfo.PUT("/",s.ModifySelfDetail)
    }
    work:=r.Group("/work/:phone")
    {
        work.POST("/",s.MakeWorkFile)
        work.GET("/",s.ViewWorkFile)
    }
    s.Router=r
    r.Run(":8080")
}