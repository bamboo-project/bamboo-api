package service

import "bamboo-api/app/dao"

type userService struct {
	userDao *dao.UserDao
}

var UserService *userService

func InitUserService(userDao *dao.UserDao) {
	UserService = &userService{
		userDao: userDao,
	}
}

func UpdateUserInfo() {

}
