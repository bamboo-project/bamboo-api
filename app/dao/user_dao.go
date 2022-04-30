package dao

import (
	"bamboo-api/app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) Create(user models.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		log.Warnf("[UserDao] create user err, user:%+v, err:%+v", user, err)
	}
	return nil
}

func (u *UserDao) Get(address string) (user *models.User, err error) {
	err = u.db.Select("wallet_address = ?", address).Find(&user).Error
	if err != nil {
		log.Warnf("[UserDao] get user err, key:%+v, err:%+v", address, err)
		return nil, err
	}
	return
}

func (u *UserDao) Update() {

}
