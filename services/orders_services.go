package services

import (
	"fmt"

	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"
)

func Getordersproduct(order configs.Postoder) configs.Products {
	db := configs.GetDB()
	res := db.QueryRow("select * from products where id=?", order.Productid)
	var product configs.Products
	res.Scan(&product.Id, &product.Name, &product.Quantity, &product.Price)
	return product
}

func Checkquantity(order configs.Postoder, product configs.Products) bool {
	if order.Quantity > product.Quantity {
		fmt.Println("false")
		return false
	} else {
		return true
	}
}

func Gettotalprice(order configs.Postoder, product configs.Products) float32 {
	return product.Price * float32(order.Quantity)
}

func Createorder(order configs.Order) {
	db := configs.GetDB()
	db.Exec("insert into orders(productid,quantity,price,location) values($1,$2,$3,$4)", order.Productid, order.Quantity, order.Price, order.Location)
	db.Exec("update products set quantity=quantity-$1 where id=$2", order.Quantity, order.Productid)

}
