package main

import (
	"fmt"
	"log"
	"port-adapter/port"
	"port-adapter/repository/mongodb"
	"port-adapter/repository/mysql"
)

func main() {
	var repo port.UserRespository
	useDb := "mysql"

	switch useDb {
	case "mysql":
		repo = mysql.NewUserRespositoryMysqlImpl()
	case "mongodb":
		repo = mongodb.NewUserRespositoryMongodbImpl()
	}

	err := GetUser(repo)
	log.Println(err)
}

func GetUser(repo port.UserRespository) error {
	user, err := repo.FindUserById(1)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return nil
}
