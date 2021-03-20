package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)
//登陆验证
func (s *Service) Login (c *gin.Context)  {
    input := new(AccountTable)
    err:=c.BindJSON(input)
    if err!=nil{
        c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"can not bind json"))
    }
    fmt.Println("账号密码； ",input)
    err=s.DB.Where("account_phone=?","password=?",input.AccountPhone,input.Password).Find(input).Error
    if err!=nil{
        fmt.Println("没有这个人")
        c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"can not find this username"))
    }
    c.JSON(MakeSuccessReturn(200))
}
//注册
func (s Service) Register( c *gin.Context)  {
    register:= new(AccountTable)
    err:=c.ShouldBind(register)
    fmt.Println(*register)
    if err!=nil{
        c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"can not bind json"))
    }
    if register.AccountPhone == "" || register.Password == "" {
        c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"invalid data"))
        return
    }
        tx:=s.DB.Begin()
        {
           if result:=s.DB.Create(&AccountTable {
               AccountPhone: register.AccountPhone,
               Password: register.Password,
               DingdingAccount: register.DingdingAccount,
           });result.Error!=nil||result.RowsAffected>=3{
               fmt.Println("数据库错误",err)
               tx.Rollback()
               c.JSON(MakeErrorReturn(http.StatusBadRequest,40000,"register fail"))
               return
           }
           c.JSON(MakeSuccessReturn(200))
        }
}
