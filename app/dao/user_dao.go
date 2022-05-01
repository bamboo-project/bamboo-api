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

func (u *UserDao) Create(user *models.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		log.Warnf("[UserDao] create user err, user:%+v, err:%+v", user, err)
	}
	return nil
}

func (u *UserDao) Get(address string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("wallet_address = ?", address).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Warnf("[UserDao] get user err, key:%+v, err:%+v", address, err)
		return nil, err
	}
	return user, nil
}

func (u *UserDao) Update(user *models.User) error {
	err := u.db.Updates(user).Error
	if err != nil {
		log.Warnf("[UserDao] update user err, user:%+v, err:%+v", user, err)
		return err
	}
	return nil
}
