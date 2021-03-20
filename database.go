package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

func (s *Service) initDB()  {
    var dbConf= s.Conf.Mysql
    //link mysql
    dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbConf.User,
        dbConf.Password,
        dbConf.Addr,
        dbConf.DBname,
    )
    fmt.Println(dsn)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
    if err!=nil{
        panic("连接数据库出错")
    }
    fmt.Println("[-] 数据库连接成功")
    s.DB = db
}
