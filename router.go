package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)
//为了解决跨域
func cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
        c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT ")
        c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
        c.Header("Access-Control-Allow-Credentials", "true")
        if method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
        }
        c.Next()
    }
}
func (s Service)initRouter()  {
    r:=gin.Default()
    r.Use(cors())
    //注册登陆模块
    r.POST("/login",s.Login)
    r.POST("/register",s.Register)
    selfInfo:=r.Group("/profiles/:phone")
    {
       selfInfo.GET("",s.GetSelfDetail)
       selfInfo.PUT("",s.ModifySelfDetail)
    }
    work:=r.Group("/work/:phone")
    {
        work.POST("",s.MakeWorkFile)
        work.GET("",s.ViewWorkFile)
        work.PUT("",s.PromotionPost)
    }

    s.Router=r
    r.Run(":8080")
}