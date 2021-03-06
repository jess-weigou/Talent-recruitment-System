package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-basic/uuid"
    "time"
)
//login
func (s *Service) Login (c *gin.Context)  {
    input := new(AccountTable)
    err:=c.ShouldBind(input)
    if err!=nil{
        c.JSON(MakeErrorReturn("can not bind json"))
    }
    input.Password = Md5Encryption(input.Password)
    s.DB.Where(&AccountTable{
        AccountPhone: input.AccountPhone,
        Password: input.Password,
    }).Find(&input)
    if input.Position == ""{
        fmt.Println("账号或者密码错误")
        c.JSON(MakeErrorReturn("username or password wrong"))
        return
    }
    //get uuid and insert to mysql
    uuid:=uuid.New()
    s.DB.Where(&AccountTable{
        AccountPhone:input.AccountPhone,
    }).Updates(&AccountTable{
        Uuid: uuid,
    })
    fmt.Println(uuid)
    c.JSON(MakeSuccessReturn(uuid))
}
//register
func (s Service) Register( c *gin.Context)  {
    register:= new(AccountTable)
    err:=c.ShouldBind(register)
    fmt.Println("注册信息",*register)
    if err!=nil{
        c.JSON(MakeErrorReturn("can not bind json"))
    }
    if register.AccountPhone == "" || register.Password == "" {
        c.JSON(MakeErrorReturn("invalid data"))
        return
    }

    s.DB.Where(&AccountTable{
            AccountPhone: register.AccountPhone,
    }).Find(&register)
    if register.Position!=""{
        c.JSON(MakeErrorReturn("the user has been registered"))
        return
    }
    //加密
    register.Password = Md5Encryption(register.Password)

    //insert to table of AccountTable、StaffInterface、SelfDetails
            s.DatabaseCommit(&AccountTable {
                AccountPhone: register.AccountPhone,
                Password: register.Password,
                DingdingAccount: register.DingdingAccount,
                Position: "1",  //注册默认为员工
            },c,"register fail")
            s.DatabaseCommit(&StaffInterface{
                StaffPhone:register.AccountPhone,
            },c,"register fail_StaffInterface")
            s.DatabaseCommit(&SelfDetails{
               StaffPhone2:register.AccountPhone,
            },c,"register fail_SelfDetails")
           c.JSON(MakeSuccessReturn(""))
}
//modify selfDetails
func (s Service)ModifySelfDetail (c *gin.Context) {
   selfInformation:=new(SelfDetails)
    phone:=c.Param("phone")
    s.DB.Where(&SelfDetails{
        StaffPhone2: phone,
    }).Find(&selfInformation)
    //s.DB.Where("staff_phone2=?",phone).Find(&selfInformation)
    if selfInformation.StaffPhone2!=""{
        err := c.ShouldBind(selfInformation)
        if err != nil {
            fmt.Println(err)
            c.JSON(MakeErrorReturn("can not bind"))
            return
        }
        //s.DB.Debug().Where("staff_phone2=?",phone).Updates(&selfInformation)
        s.DB.Debug().Where(&SelfDetails{
                StaffPhone2: phone,
            }).Updates(&selfInformation)
        c.JSON(MakeSuccessReturn(""))
    }else{
        c.JSON(MakeErrorReturn("can not find this people"))
    }
}
//view selfDetails
func (s Service)GetSelfDetail(c *gin.Context)  {
    selfInformation := new(SelfDetails)
    phone:=c.Param("phone")
    //phone = s.Encryption(phone)
    fmt.Println("接受到的电话号码是：",phone)
    s.DB.Where(&SelfDetails{
       StaffPhone2: phone,
    }).Find(&selfInformation)
    //s.DB.Debug().Where("staff_phone2=?",phone).Find(&selfInformation)
    fmt.Println(selfInformation)
    if selfInformation.StaffPhone2==""{
        c.JSON(MakeErrorReturn("can not find this people"))
        return
    }
    //phone = s.Decryption(phone)
    c.JSON(MakeSuccessReturn(selfInformation))
}
//make the work file
func (s Service)MakeWorkFile(c *gin.Context)  {
    authorization:=c.GetHeader("Authorization")
    fmt.Println(authorization)
    account:=new(AccountTable)
    phone:=c.Param("phone")
    fmt.Println("work_file",phone)
    //acknowledge the user's Auth
    s.DB.Where(AccountTable{
        Uuid: authorization,
    }).Find(&account)
    //s.DB.Where("uuid=?",authorization).Find(&account)
    if account.Position==""{
        fmt.Println("can not find this people")
        c.JSON(MakeErrorReturn("can not find this people"))
        return
    }else if account.Position<"2"{
        fmt.Println("You don't have authority ")
       c.JSON(MakeErrorReturn("You don't have authority "))
       return
    }
    workFile:=new(EmploymentStatus)
    err:=c.ShouldBind(workFile)
    workFile.WorkInTime=time.Now()
    if err != nil {
        fmt.Println("can not bind data")
        c.JSON(MakeErrorReturn("can not bind data"))
        return
    }
    fmt.Println(workFile)
    //实际使用时没有这个
    //s.DatabaseCommit(&CompanyInterface{
    //    CompanyId: workFile.CompanyId,
    //},c,"fail_CompanyId")
    workFile.StaffPhone1 = phone
    s.DatabaseCommit(&workFile,c,"can not insert_fail_CompanyId")
    c.JSON(MakeSuccessReturn(""))
}
//view the work file
func (s Service)ViewWorkFile(c *gin.Context) {
    authorization:=c.GetHeader("Authorization")
    fmt.Println(authorization)
    phone:=c.Param("phone")
    fmt.Println("接受到的手机号：",phone)
    workFile:=new(EmploymentStatus)
    s.DB.Where(EmploymentStatus{
        StaffPhone1: phone,
    }).Find(&workFile)
    //s.DB.Where("staff_phone1=?",phone).Find(&workFile)
    fmt.Println("workFile",workFile)
    if workFile.CompanyId==""{
        fmt.Println("can not find this user")
        c.JSON(MakeErrorReturn("can not find this user"))
        return
    }else{
        c.JSON(MakeSuccessReturn(workFile))
    }
}
//promotion post (提升职位)
func (s Service)PromotionPost(c *gin.Context)  {
    //测试用
    acc :=new(AccountTable)
    err:=c.ShouldBind(&acc)
    fmt.Println("接受到的职位:",acc.Position)
    if err!=nil{
       c.JSON(MakeErrorReturn("can't bind json"))
       return
    }
    authorization:=c.GetHeader("Authorization")
    accountBoss:= new(AccountTable)
    accountStaff:=new(AccountTable)

    //find boss and staff
    phone := c.Param("phone")

    s.DB.Where(&AccountTable{
        Uuid: authorization,
    }).Find(&accountBoss)
    //s.DB.Where("uuid=?",authorization).Find(&accountBoss)
    if accountBoss.AccountPhone==""{
        c.JSON(MakeErrorReturn("can't find this head of department"))
        return
    }
    s.DB.Where(&AccountTable{
        AccountPhone: phone,
    }).Find(&accountStaff)
    //s.DB.Where("account_phone=?",phone).Find(&accountStaff)
    if accountStaff.AccountPhone==""{
        c.JSON(MakeErrorReturn("can't find this staff"))
        return
    }
    fmt.Println("部门主管信息",accountBoss)
    if accountBoss.Position<="2"{
        c.JSON(MakeErrorReturn("you don't have authority"))
        return
    }else if accountBoss.Position<="4"{
        accountStaff.Position = acc.Position
        s.DB.Where(&AccountTable{
            AccountPhone: phone,
        }).Updates(&accountStaff)
        //s.DB.Where("account_phone=?",phone).Updates(&accountStaff)
        c.JSON(MakeSuccessReturn(""))
    }else{
        c.JSON(MakeErrorReturn("unexpected error"))
        return
    }
}
func (s Service)ClearPhone(c *gin.Context)  {
    phone := c.Param("phone")
    staffInterface := new(StaffInterface)
    accountTable := new(AccountTable)
    selfDetails :=new(SelfDetails)
    employmentStatus :=new(EmploymentStatus)
    s.DB.Where(&EmploymentStatus{
        StaffPhone1: phone,
    }).Delete(&employmentStatus)
    s.DB.Where(&StaffInterface{
        StaffPhone: phone,
    }).Delete(&staffInterface)

    s.DB.Where(&AccountTable{
        AccountPhone: phone,
    }).Delete(&accountTable)
    s.DB.Where(&SelfDetails{
        StaffPhone2: phone,
    }).Delete(&selfDetails)
    c.JSON(MakeSuccessReturn(""))
}