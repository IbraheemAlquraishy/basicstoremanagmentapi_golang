package configs

type Products struct {
	Id       int
	Name     string
	Quantity int
	Price    float32
	img      string
}

type Order struct {
	Id        int
	Productid int
	Quantity  int
	Price     float32
	Location  string
}

type Postoder struct {
	Productid int
	Quantity  int
	Location  string
}

type Login struct {
	Name     string
	Password string
}
