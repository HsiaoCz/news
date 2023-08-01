package service

import "github.com/gin-gonic/gin"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) UserRegister(c *gin.Context) {}
