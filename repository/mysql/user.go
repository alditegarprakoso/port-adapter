package mysql

import (
	"port-adapter/entity"

	"gorm.io/gorm"
)

type UserRespositoryMysqlImpl struct {
	db *gorm.DB
}

func NewUserRespositoryMysqlImpl(db *gorm.DB) *UserRespositoryMysqlImpl {
	println("MYSQL USER DIPAKAI")
	return &UserRespositoryMysqlImpl{db: db}
}

func (u *UserRespositoryMysqlImpl) CreateUser(user *entity.User) (*entity.User, error) {
	result := u.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserRespositoryMysqlImpl) FindUserById(id int) (*entity.User, error) {
	user := &entity.User{}
	result := u.db.First(user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserRespositoryMysqlImpl) FindUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}
