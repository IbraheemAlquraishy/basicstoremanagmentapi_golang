package services

import "github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"

func Checklog(log configs.Login) bool {
	db := configs.GetDB()
	res := db.QueryRow("select * from admins where name=$1 and password=$2")
}
