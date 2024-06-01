package main

import (
	"fmt"
	"log"
	"port-adapter/infrastructure"
	"port-adapter/port"
	"port-adapter/repository/mysql"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func main() {
	db := infrastructure.NewGorm()
	repo := mysql.NewUserRespositoryMysqlImpl(db)
	err := Execute(repo)
	log.Println(err)
}

func Execute(repo port.UserRespository) error {
	// user := &entity.User{
	// 	Name: "Aldi Tegar Prakoso",
	// 	Age:  25,
	// }

	// user, err := repo.CreateUser(user)

	// if err != nil {
	// 	return err
	// }

	// fmt.Println(user)

	user, err := repo.FindUserById(1)

	if err != nil {
		return err
	}

	fmt.Println(user)

	return nil
}
