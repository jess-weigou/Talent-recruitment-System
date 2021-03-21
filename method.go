package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-basic/uuid"
)
//登陆验证
func (s *Service) Login (c *gin.Context)  {
    input := new(AccountTable)
    err:=c.ShouldBind(input)
    if err!=nil{
        c.JSON(MakeErrorReturn("can not bind json"))
    }
    fmt.Println("账号密码； ",input)
    s.DB.Where("account_phone=?","password=?",input.AccountPhone,input.Password).Find(&input)
    if input.DingdingAccount == ""{

        fmt.Println("没有这个人")
        c.JSON(MakeErrorReturn("can not find this username"))
        return
    }
    //get uuid and insert to mysql
    uuid:=uuid.New()
    s.DB.Where("account_phone=?",input.AccountPhone).Updates(&AccountTable{
        Uuid: uuid,
    })
    fmt.Println(uuid)
    c.JSON(MakeSuccessReturn(uuid))
}
//注册
func (s Service) Register( c *gin.Context)  {
    register:= new(AccountTable)
    err:=c.ShouldBind(register)
    fmt.Println(*register)
    if err!=nil{
        c.JSON(MakeErrorReturn("can not bind json"))
    }
    if register.AccountPhone == "" || register.Password == "" {
        c.JSON(MakeErrorReturn("invalid data"))
        return
    }
        tx:=s.DB.Begin()
        {
           if s.DB.Create(&AccountTable {
               AccountPhone: register.AccountPhone,
               Password: register.Password,
               DingdingAccount: register.DingdingAccount,
           }).RowsAffected!=1{
               fmt.Println("数据库错误",err)
               tx.Rollback()
               c.JSON(MakeErrorReturn("register fail"))
               return
           }
           tx.Commit()
            uuid:=uuid.New()
            s.DB.Where("account_phone=?",register.AccountPhone).Updates(&AccountTable{
                Uuid: uuid,
            })
           c.JSON(MakeSuccessReturn(uuid))
        }
}
//增加个人信息
func (s Service)AddSelfDetail (c *gin.Context) {
    selfInformation := new(SelfDetails)
    err := c.ShouldBind(selfInformation)
    if err != nil {
        MakeErrorReturn("can not bind json")
        return
    }
    tx := s.DB.Begin()
    {
        if result := s.DB.Create(&selfInformation); result.Error != nil || result.RowsAffected >= 3 {
            fmt.Println(result.Error,"数据库错误")
            tx.Rollback()
            c.JSON(MakeErrorReturn("add self_detail fail"))
            return
        }
        tx.Commit()
        c.JSON(MakeSuccessReturn(""))
    }
}
