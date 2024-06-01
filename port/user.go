package port

import "port-adapter/entity"

type UserRespository interface {
	FindUserById(id int) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}
