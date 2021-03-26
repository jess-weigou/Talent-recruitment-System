package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-basic/uuid"
    "reflect"
)
//登陆验证
func (s *Service) Login (c *gin.Context)  {
    input := new(AccountTable)
    err:=c.ShouldBind(input)
    if err!=nil{
        c.JSON(MakeErrorReturn("can not bind json"))
    }
    fmt.Println("账号密码； ",input.AccountPhone,input.Password)
    s.DB.Debug().Select("position").Where("account_phone=? AND password=?",input.AccountPhone,input.Password).Find(&input)
    if input.Position == ""{
        fmt.Println("账号或者密码错误")
        c.JSON(MakeErrorReturn("username or password wrong"))
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
    s.DB.Select("password").Where("account_phone=?",register.AccountPhone).Find(&register)
    if register.Position!=""{
        c.JSON(MakeErrorReturn("用户已注册"))
        return
    }
    //注册插入各表电话信息
            s.DatabaseCommit(&AccountTable {
                AccountPhone: register.AccountPhone,
                Password: register.Password,
                DingdingAccount: register.DingdingAccount,
                Position: "1",  //注册默认为员工
            },c,"register fail")
            s.DatabaseCommit(&StaffInterface{
                StaffPhone:register.AccountPhone,
            },c,"register fail_StaffInterface")
            //s.DatabaseCommit(&SelfDetails{
            //    StaffPhone2:register.AccountPhone,
            //},c,"register fail_SelfDetails")
            uuid:=uuid.New()
            s.DB.Where("account_phone=?",register.AccountPhone).Updates(&AccountTable{
                Uuid: uuid,
            })
           c.JSON(MakeSuccessReturn(uuid))
}
//修改个人信息
func (s Service)ModifySelfDetail (c *gin.Context) {
    selfInformation:=new(SelfDetails)
    err := c.ShouldBind(selfInformation)
    if err != nil {
       MakeErrorReturn("can not bind")
       return
    }
    var typeInfo = reflect.TypeOf(*selfInformation)
    var valInfo = reflect.ValueOf(*selfInformation)
    phone:=c.Param("phone")
    num := typeInfo.NumField()
    for i:=0;i<num;i++{
        if valInfo.Field(i).String()+""!=""{
            s.DB.Where("staff_phone2=?",phone).Updates(SelfDetails{
                typeInfo.Field(i).Name:valInfo.Field(i).FieldByName(typeInfo.Field(i).Name),
            })
        }
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
