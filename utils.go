package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type SuccessReturn struct {
    Status bool `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  interface{}   `json:"data"`
}
type ErrorReturn struct {
    Status bool    `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  interface{}   `json:"data"`
}

func MakeSuccessReturn(data interface{})(int ,interface{})  {
    return 200,SuccessReturn{
        Status: true,
        Msg: "success",
        TokenData: data,
    }
}
func MakeErrorReturn(msg string)(int ,interface{})  {
    return 400,ErrorReturn{
        Status: false,
        Msg: msg,
        TokenData: "",
    }
}
func (s Service)DatabaseCommit(data interface{},c *gin.Context,msg string)  {
    tx:=s.DB.Begin()
    {
        if s.DB.Create(data).RowsAffected!=1{
            fmt.Println(msg)
            tx.Rollback()
            c.JSON(MakeErrorReturn(msg))
        }
        tx.Commit()

    }
}