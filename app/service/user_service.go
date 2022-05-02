package service

import (
	"bamboo-api/app/constans"
	"bamboo-api/app/dao"
	"bamboo-api/app/models"
	"bamboo-api/app/pkg/entity/dto"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type userService struct {
	userDao *dao.UserDao
}

var UserService *userService

func InitUserService(userDao *dao.UserDao) {
	UserService = &userService{
		userDao: userDao,
	}
}
func (u *userService) InitUser(walletId string) error {
	err := u.userDao.Create(&models.User{
		WalletAddress: walletId,
		WalletType:    constans.NeoWallet,
		AvatarUrl:     "",
		BambooCoins:   0,
		Followers:     0,
		Following:     0,
		IsTwitter:     0,
	})

	return err
}

func (u *userService) GetUser(walletId string) (*dto.Response, error) {
	user, err := u.userDao.Get(walletId)
	if nil != err {
		log.Errorf("[BindTwitter] get user failed, user=%v, err=%+v", user, err)
		return nil, errors.New("bind twitter failed")
	}
	if nil == user {
		err = u.InitUser(walletId)
		if nil == err {
			user, err = u.userDao.Get(walletId)
		}
	}
	return &dto.Response{
		Code: 0,
		Data: user,
	}, nil
}

func UpdateUserInfo() {

}

func (u *userService) BindTwitter(walletAddress string, twitterId string, twitterUrl string, twitterToken string, twitterSecret string) error {
	user, err := u.userDao.Get(walletAddress)
	if nil != err || nil == user {
		log.Errorf("[BindTwitter] get user failed, user=%v, err=%+v", user, err)
		return errors.New("bind twitter failed")
	}
	user.IsTwitter = 1
	user.TwitterId = twitterId
	user.TwitterUrl = twitterUrl
	user.TwitterToken = twitterToken
	user.TwitterSecret = twitterSecret
	err = u.userDao.Update(user)
	return nil
}
