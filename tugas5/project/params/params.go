package params

import (
	"errors"
	"tugas5/project/model"

	"github.com/go-playground/validator/v10"
)

type AuthRegister struct {
	Email    string
	Password string
}

type UserCreate struct {
	Fullname   string
	Gender     string
	Contact    string
	Street     string
	CityId     string
	ProvinceId string
}

type UserLogin struct {
	Email    string
	Password string
}

type CartCreate struct {
	UserId    string
	ProductId string
	Qty       uint16
	Subtotal  uint32
}

type CategoryCreate struct {
	Name        string
	Description string
	UserId      string
}

type ProductCreate struct {
	Name        string
	Description string
	Price       uint32
	Stock       uint16
	CategoryId  string
	UserId      string
}

type CartUpdate struct {
	CartCreate
}

type CategoryUpdate struct {
	CategoryCreate
}

type ProductUpdate struct {
	ProductCreate
}

type UserUpdate struct {
	UserCreate
}

func (u *AuthRegister) ParseToModelRegister() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserCreate) ParseToModelCreate() *model.User {
	return &model.User{
		Fullname:   u.Fullname,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}
}

func Validate(u interface{}) error {
	err := validator.New().Struct(u)
	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
}

func (p *CartCreate) ParseToModel() *model.Cart {
	return &model.Cart{
		ProductId: p.ProductId,
		Qty:       p.Qty,
	}
}

func (c *CategoryCreate) ParseToModel() *model.Category {
	return &model.Category{
		Name:        c.Name,
		Description: c.Description,
		UserId:      c.UserId,
	}
}

func (p *ProductCreate) ParseToModel() *model.Product {
	return &model.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		UserId:      p.UserId,
	}
}
