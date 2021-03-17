package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Service struct {
    Conf *Config
    DB  *gorm.DB
    Router *gin.Engine
}

func (s *Service) init() {
    s.initConfig()
    s.initDB()
}