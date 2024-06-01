package mysql

import "port-adapter/entity"

type UserRespositoryMysqlImpl struct{}

func NewUserRespositoryMysqlImpl() *UserRespositoryMysqlImpl {
	println("MONGODB USER DIPAKAI")
	return &UserRespositoryMysqlImpl{}
}

func (u *UserRespositoryMysqlImpl) FindUserById(id int) (*entity.User, error) {
	return &entity.User{
		Name: "Aldi Tegar Prakoso",
		Age:  23,
	}, nil
}

func (u *UserRespositoryMysqlImpl) FindUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}
