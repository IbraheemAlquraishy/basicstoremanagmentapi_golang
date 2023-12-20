package configs

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

func Gentoken(admin Admin) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  admin.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedtoken, err := token.SignedString(jwtKey)
	fmt.Println(err)

	return signedtoken
}

func Verifytoken(tokens string) (*Admin, error) {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		fmt.Println(claims["id"], claims["exp"])
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			return nil, errors.New("expired token")
		}
		adminid := claims["id"].(float64)
		fmt.Println(adminid)
		admin, err := Getadminbyid(int64(adminid))
		if err != nil {
			return nil, err
		}
		return admin, nil

	} else {
		return nil, errors.New("invalid token")
	}

}

func Getadminbyid(id int64) (*Admin, error) {
	db := GetDB()
	var admin Admin
	db.QueryRow("select * from admin where id=?", id).Scan(&admin.Id, &admin.Name, &admin.Password)
	if admin.Id == 0 {
		return nil, errors.New("not such admin")
	}
	return &admin, nil

}
