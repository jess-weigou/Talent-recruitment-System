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
            s.DatabaseCommit(&AccountTable {
                AccountPhone: register.AccountPhone,
                Password: register.Password,
                DingdingAccount: register.DingdingAccount,
            },c)
           if s.DB.Create().RowsAffected!=1{
               fmt.Println("数据库错误",err)
               tx.Rollback()
               c.JSON(MakeErrorReturn("register fail"))
               return
           }else{
               s.DB.Create(&EmploymentStatus{
                   StaffPhone1: register.AccountPhone,
               })
               s.DB.Create(SelfDetails{
                   StaffPhone2:register.AccountPhone,
               })
               s.DB.Create(StaffInterface{
                   StaffPhone:register.AccountPhone,
               })
           }
           tx.Commit()

            uuid:=uuid.New()
            s.DB.Where("account_phone=?",register.AccountPhone).Updates(&AccountTable{
                Uuid: uuid,
            })
           c.JSON(MakeSuccessReturn(uuid))
        }
}
//修改个人信息
func (s Service)ModifySelfDetail (c *gin.Context) {
    selfInformation := new(SelfDetails)
    err := c.ShouldBind(selfInformation)
    if err != nil {
        MakeErrorReturn("can not bind json")
        return
    }
    phone:=c.Param("phone")

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
//查看个人信息
func (s Service)GetSelfDetail(c *gin.Context)  {
    var Authorization string
    selfInformation := new(SelfDetails)
    err:=c.ShouldBind(Authorization)
    if err != nil {
        MakeErrorReturn("can not bind json")
        return
    }
    phone:=c.Param("phone")
    s.DB.Where("staff_phone2=?",phone).Find(selfInformation)
    if selfInformation.StaffPhone2==""{
        MakeErrorReturn("数据库错误")
        return
    }
    MakeSuccessReturn(selfInformation)
}
