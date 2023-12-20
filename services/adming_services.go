package services

import (
	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"

	"golang.org/x/crypto/bcrypt"
)

func Checklog(log configs.Login) (configs.Admin, error) {
	db := configs.GetDB()
	var admin configs.Admin

	res := db.QueryRow("select * from admin where name=$1 ", log.Name)
	res.Scan(&admin.Id, &admin.Name, &admin.Password)

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(log.Password))
	return admin, err
}

func Insertadmin(log configs.Login) error {
	db := configs.GetDB()
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(log.Password), bcrypt.MinCost)

	if err != nil {
		return err
	}
	_, err = db.Exec("insert into admin(name,password) values($1,$2)", log.Name, hashedPasswordBytes)
	return err
}
