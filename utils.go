package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type Data struct {
    Token interface{} `json:"token"`
}
type SuccessReturn struct {
    Status bool `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  Data   `json:"data"`
}
type ErrorReturn struct {
    Status bool    `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  Data   `json:"data"`
}

func MakeSuccessReturn(data interface{})(int ,interface{})  {
    return 200,SuccessReturn{
        Status: true,
        Msg: "success",
        TokenData: Data{Token: data},
    }
}
func MakeErrorReturn(msg string)(int ,interface{})  {
    return 200,ErrorReturn{
        Status: false,
        Msg: msg,
        TokenData: Data{Token: ""},
    }
}
func (s Service)DatabaseCommit(data interface{},c *gin.Context)  {
    tx:=s.DB.Begin()
    {
        if s.DB.Create(&data).RowsAffected!=1{
            fmt.Println("数据库错误")
            tx.Rollback()
            c.JSON(MakeErrorReturn("database fail"))
            return
        }
        tx.Commit()
        
    }
}