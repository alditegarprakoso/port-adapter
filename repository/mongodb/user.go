package mongodb

import "port-adapter/entity"

type UserRespositoryMongodbImpl struct{}

func NewUserRespositoryMongodbImpl() *UserRespositoryMongodbImpl {
	println("MONGODB USER DIPAKAI")
	return &UserRespositoryMongodbImpl{}
}

func (u *UserRespositoryMongodbImpl) FindUserById(id int) (*entity.User, error) {
	return nil, nil
}

func (u *UserRespositoryMongodbImpl) FindUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}
