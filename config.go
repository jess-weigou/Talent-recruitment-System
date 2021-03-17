package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
)

type Mysql struct {
    User string
    Password string
    Addr string
    DBname string
}
type Config struct {
    Mysql Mysql
}
func (s *Service) initConfig()  {
    c := new(Config)
    _ , err:=toml.DecodeFile("./config_talent_system/config.toml",&c)
    if err!=nil{
        panic("加载配置文件出错")
    }
    s.Conf = c
    fmt.Println("[-] 加载配置文件成功")
}
