package respone

import (
	"net/http"
	"tugas5/project/model"
)

type UserCreateResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserFindAllResponse struct {
	Id         string `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
}

type UserUpdateResponse struct {
	Id         uint   `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Info    string      `json:"info"`
	Payload interface{} `json:"payload"`
	Error   interface{} `json:"error"`
}

type CartCreateResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

type CartUpdateResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

type CartGetResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

type CategoryCreateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUpdateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryGetAllResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductCreateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Description string `json:"description"`
	Stock       uint16 `json:"stock"`
}

type ProductUpdateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Description string `json:"description"`
	Stock       uint16 `json:"stock"`
}

type ProductGetAllResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"namdescription"`
	Price       uint32 `json:"price"`
	Stock       uint16 `json:"stock"`
}

func NewUserCreateResponse(user *model.User) *UserCreateResponse {
	return &UserCreateResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
}

func NewUserFindAllResponse(users *[]model.User) []UserFindAllResponse {
	var usersFind []UserFindAllResponse
	for _, user := range *users {
		usersFind = append(usersFind, *parseModelToUserFind(&user))
	}
	return usersFind
}

func NewUserUpdateResponse(user *model.User) *UserUpdateResponse {
	return &UserUpdateResponse{
		Id:         user.ID,
		Fullname:   user.Fullname,
		Email:      user.Email,
		Gender:     user.Gender,
		Contact:    user.Contact,
		Street:     user.Street,
		CityId:     user.CityId,
		ProvinceId: user.ProvinceId,
	}
}
func parseModelToUserFind(user *model.User) *UserFindAllResponse {
	return &UserFindAllResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Gender:   user.Gender,
		Contact:  user.Contact,
		Street:   user.Street}
}

func SuccessCreated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: "success create",
		Payload: payload,
	}
}

func SuccessUpdated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success updated",
		Payload: payload,
	}
}

func SuccessDeleted(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success delete",
		Payload: payload,
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success get data",
		Payload: payload,
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Error:  err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err,
	}
}
func ErrNotFound() *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "Data not found",
		Error:   "NO_DATA",
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   "Unauthorized",
	}
}

func NewCartCreateResponse(cart *model.Cart) *CartCreateResponse {
	return &CartCreateResponse{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Qty:       cart.Qty,
		Subtotal:  cart.Subtotal,
	}
}

func NewCartUpdateResponse(cart *model.Cart) *CartUpdateResponse {
	return &CartUpdateResponse{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Qty:       cart.Qty,
		Subtotal:  cart.Subtotal,
	}
}

func NewCartGetResponse(carts *[]model.Cart) *[]CartGetResponse {
	var cartsResponse []CartGetResponse

	for _, cart := range *carts {
		cartsResponse = append(cartsResponse, CartGetResponse{
			Id:        cart.ID,
			UserId:    cart.UserId,
			ProductId: cart.ProductId,
			Qty:       cart.Qty,
			Subtotal:  cart.Subtotal,
		})
	}

	return &cartsResponse
}

func NewCategoryCreateResponse(category *model.Category) *CategoryCreateResponse {
	return &CategoryCreateResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func NewCategoryUpdateResponse(category *model.Category) *CategoryUpdateResponse {
	return &CategoryUpdateResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func NewCategoryGetAllResponse(categories *[]model.Category) *[]CategoryGetAllResponse {
	var categoriesResponse []CategoryGetAllResponse

	for _, category := range *categories {
		categoriesResponse = append(categoriesResponse, CategoryGetAllResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &categoriesResponse
}

func NewProductCreateResponse(product *model.Product) *ProductCreateResponse {
	return &ProductCreateResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func NewProductUpdateResponse(product *model.Product) *ProductUpdateResponse {
	return &ProductUpdateResponse{
		Id:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
	}
}

func NewProductGetAllResponse(products *[]model.Product) *[]ProductGetAllResponse {
	var productsResponse []ProductGetAllResponse

	for _, product := range *products {
		productsResponse = append(productsResponse, ProductGetAllResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		})
	}

	return &productsResponse
}
