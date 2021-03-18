package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func (s *Service) Login (c *gin.Context)  {
    input := new(Input)
    err:=c.BindJSON(input)
    if err!=nil{
        c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"can not bind json"))
    }
    fmt.Println("账号密码； ",input)

    c.JSON(MakeSuccessReturn(200))
}
