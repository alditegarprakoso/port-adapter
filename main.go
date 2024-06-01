package main

import (
	"adnin-port-adapter/port"
	"adnin-port-adapter/repository/mongodb"
	"adnin-port-adapter/repository/mysql"
	"fmt"
	"log"
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
