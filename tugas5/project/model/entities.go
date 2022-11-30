package model

import "gorm.io/gorm"

var Users = []User{}
var Categories = []Category{}
var Carts = []Cart{}
var OrderItems = []OrderItem{}
var Orders = []Order{}
var Payments = []Payment{}
var Deliveries = []Delivery{}

type User struct {
	gorm.Model
	Fullname   string
	Email      string
	Password   string
	Role       string
	Gender     string
	Contact    string
	Street     string
	CityId     string
	ProvinceId string
}

type Cart struct {
	gorm.Model
	UserId    string
	ProductId string
	Qty       uint16
	Subtotal  uint32
	User      User
	Product   Product
}

type Category struct {
	gorm.Model
	Name        string
	Description string
	UserId      string
	User        User
}

type Product struct {
	gorm.Model
	Name        string
	Price       uint32
	Stock       uint16
	Description string
	CategoryId  string
	UserId      string
	Category    Category
	User        User
}

type OrderItem struct {
	gorm.Model
	OrderId   string
	ProductId string
	Qty       uint16
	SubTotal  uint32
	Order     Order
	Product   Product
}

type Order struct {
	gorm.Model
	UserId string
	Total  uint32

	OrderItem []OrderItem
	Payment   []Payment
	User      User
}

type Payment struct {
	gorm.Model
	OrderId  string
	Amount   uint32
	Provider string
	Status   string
	Order    Order
}

type Delivery struct {
	gorm.Model
	OrderId       string
	ReceiptNumber string
	Expedition    string
	Weight        uint32
	ServiceType   string
	Cost          uint32
	Estimation    string
	Order         Order
}
