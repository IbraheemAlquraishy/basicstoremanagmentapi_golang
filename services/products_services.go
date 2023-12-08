package services

import (
	"fmt"

	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"
)

func Queryallproducts() []configs.Products {
	p := []configs.Products{}
	db := configs.GetDB()
	d, _ := db.Query("select * from products")

	for d.Next() {
		var temp configs.Products
		d.Scan(&temp.Id, &temp.Name, &temp.Quantity, &temp.Price)
		fmt.Print(d)
		p = append(p, temp)
	}
	return p
}
