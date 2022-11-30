package service

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	database "tugas5/project/database"
	"tugas5/project/params"
	respone "tugas5/project/respone"
	"tugas5/tokenjwt"
)

type AuthServices struct {
	face database.Userface
}

type CartService struct {
	face        database.Cartface
	productface database.Productface
	userface    database.Userface
}

func NewAuthServices(face database.Userface) *AuthServices {
	return &AuthServices{
		face: face,
	}
}

func (u *AuthServices) Register(req *params.AuthRegister) *respone.Response {
	user := req.ParseToModelRegister()

	user.Role = "user"

	hash, err := tokenjwt.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("error password %v\n", "")
		return respone.ErrInternalServer(err.Error())
	}

	user.Password = hash

	err = u.face.Create(user)
	if err != nil {
		log.Printf("error register %v\n", "")
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(user)
}

func (u *AuthServices) Login(req *params.UserLogin) *respone.Response {
	user, err := u.face.FindByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	err = tokenjwt.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return respone.ErrUnauthorized()
	}

	token := tokenjwt.Token{
		Email: user.Email,
	}

	tokString, err := tokenjwt.CreateToken(&token)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(tokString)
}

type UserServices struct {
	face database.Userface
}

func NewUserServices(face database.Userface) *UserServices {
	return &UserServices{
		face: face,
	}
}

func (u *UserServices) Create(req *params.UserCreate) *respone.Response {
	user := req.ParseToModelCreate()

	user.Role = "user"

	err := u.face.Create(user)
	if err != nil {
		log.Printf("error register %v\n", "")
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(user)
}

func (u *UserServices) FindAll() *respone.Response {

	users, err := u.face.FindAll()

	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(respone.NewUserFindAllResponse(users))
}

func (u *UserServices) FindByEmail(email string) *respone.Response {
	user, err := u.face.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrNotFound()
	}
	return respone.SuccessFindAll(user)
}

func (c *UserServices) Update(email string, req *params.UserUpdate) *respone.Response {
	user := req.ParseToModelCreate()

	getUser, err := c.face.FindByEmail(email)
	if err != nil {
		return respone.ErrBadRequest("user not found")
	}

	user.ID = getUser.ID
	err = c.face.Update(email, user)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessUpdated(respone.NewUserUpdateResponse(user))
}

type CategoryService struct {
	face database.Categoryface
}

func NewCategoryServices(face database.Categoryface) *CategoryService {
	return &CategoryService{
		face: face,
	}
}

func (c *CategoryService) GetCategories() *respone.Response {
	categories, err := c.face.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(respone.NewCategoryGetAllResponse(categories))
}

func (c *CategoryService) CreateCategory(req *params.CategoryCreate) *respone.Response {
	category := req.ParseToModel()
	fmt.Println(category)
	err := c.face.CreateCategory(category)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(respone.NewCategoryCreateResponse(category))
}

func (c *CategoryService) UpdateCategory(categoryId string, req *params.CategoryUpdate) *respone.Response {
	category := req.ParseToModel()

	getCategory, err := c.face.GetCategoryById(categoryId)
	if err != nil {
		return respone.ErrBadRequest("category not found")
	}

	category.ID = getCategory.ID
	err = c.face.UpdateCategory(categoryId, category)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessUpdated(respone.NewCategoryUpdateResponse(category))
}

func (p *CategoryService) DeleteCategory(categoryId string) *respone.Response {
	category, err := p.face.GetCategoryById(categoryId)
	if err != nil {
		return respone.ErrBadRequest("category not found")
	}

	err = p.face.DeleteCategory(categoryId, category)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessDeleted(respone.NewCategoryUpdateResponse(category))
}

func (c *CategoryService) GetCategoryById(categoryId string) *respone.Response {
	category, err := c.face.GetCategoryById(categoryId)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(category)
}

type ProductService struct {
	face         database.Productface
	categoryface database.Categoryface
}

func NewProductServices(face database.Productface, categoryface database.Categoryface) *ProductService {
	return &ProductService{
		face:         face,
		categoryface: categoryface,
	}
}

func (p *ProductService) GetProducts() *respone.Response {
	products, err := p.face.GetProducts()
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(respone.NewProductGetAllResponse(products))
}

func (p *ProductService) CreateProduct(req *params.ProductCreate) *respone.Response {
	product := req.ParseToModel()
	category, err := p.categoryface.GetCategoryById(product.CategoryId)
	if err != nil {
		return respone.ErrBadRequest("category id not found")
	}

	product.Category = *category
	err = p.face.CreateProduct(product)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(respone.NewProductCreateResponse(product))
}

func (p *ProductService) UpdateProduct(productId string, req *params.ProductUpdate) *respone.Response {
	product := req.ParseToModel()
	category, err := p.categoryface.GetCategoryById(product.CategoryId)
	if err != nil {
		return respone.ErrBadRequest("category id not found")
	}

	getProduct, err := p.face.GetProductById(productId)
	if err != nil {
		return respone.ErrBadRequest("product not found")
	}

	product.Category = *category
	product.ID = getProduct.ID
	err = p.face.UpdateProduct(productId, product)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessUpdated(respone.NewProductUpdateResponse(product))
}

func (p *ProductService) DeleteProduct(productId string) *respone.Response {
	product, err := p.face.GetProductById(productId)
	if err != nil {
		return respone.ErrBadRequest("product not found")
	}

	err = p.face.DeleteProduct(productId, product)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessDeleted(respone.NewProductUpdateResponse(product))
}

func (p *ProductService) GetProductById(productId string) *respone.Response {
	product, err := p.face.GetProductById(productId)
	if err != nil {
		return respone.ErrInternalServer(err.Error())

	}

	return respone.SuccessFindAll(product)
}

func NewCartServices(face database.Cartface, productface database.Productface, userface database.Userface) *CartService {
	return &CartService{
		face:        face,
		productface: productface,
		userface:    userface,
	}
}

func (p *CartService) GetCarts() *respone.Response {
	carts, err := p.face.GetCarts()
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(respone.NewCartGetResponse(carts))
}

func (p *CartService) GetUserCarts(userId string) *respone.Response {
	carts, err := p.face.GetUserCarts(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return respone.ErrNotFound()
		}
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(respone.NewCartGetResponse(carts))
}

func (p *CartService) GetCartById(cartId string) *respone.Response {
	cart, err := p.face.GetCartById(cartId)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessFindAll(cart)
}

func (p *CartService) CreateCart(userId string, req *params.CartCreate) *respone.Response {
	cart := req.ParseToModel()
	user, err := p.userface.FindById(userId)
	if err != nil {
		return respone.ErrBadRequest("user id not found")
	}

	product, errProduct := p.productface.GetProductById(cart.ProductId)
	if errProduct != nil {
		return respone.ErrBadRequest("product id not found")
	}

	productInCart, _ := p.face.GetCartByProductId(userId, cart.ProductId)
	if productInCart != nil {
		productInCart.Qty = productInCart.Qty + req.Qty
		productInCart.Subtotal = uint32(productInCart.Qty) * product.Price
		cartId := strconv.FormatUint(uint64(productInCart.ID), 10)
		err = p.face.UpdateCart(cartId, productInCart)
		if err != nil {
			return respone.ErrInternalServer(err.Error())
		}
		return respone.SuccessUpdated(respone.NewCartUpdateResponse(productInCart))
	}

	cart.Subtotal = product.Price * uint32(cart.Qty)
	cart.User = *user
	cart.Product = *product
	err = p.face.CreateCart(cart)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessCreated(respone.NewCartCreateResponse(cart))
}

func (p *CartService) UpdateCart(userId string, cartId string, req *params.CartUpdate) *respone.Response {
	cart := req.ParseToModel()
	user, err := p.userface.FindById(userId)
	if err != nil {
		return respone.ErrBadRequest("user id not found")
	}

	product, errProduct := p.productface.GetProductById(cart.ProductId)
	if errProduct != nil {
		return respone.ErrBadRequest("product id not found")
	}

	getCart, errCart := p.face.GetCartById(cartId)
	if errCart != nil {
		return respone.ErrBadRequest("cart not found")
	}

	cart.User = *user
	cart.Product = *product
	cart.ID = getCart.ID
	cart.Subtotal = product.Price * uint32(cart.Qty)
	err = p.face.UpdateCart(cartId, cart)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessUpdated(respone.NewCartUpdateResponse(cart))
}

func (p *CartService) DeleteCart(cartId string) *respone.Response {
	cart, err := p.face.GetCartById(cartId)
	if err != nil {
		return respone.ErrBadRequest("cart not found")
	}

	err = p.face.DeleteCart(cartId, cart)
	if err != nil {
		return respone.ErrInternalServer(err.Error())
	}

	return respone.SuccessDeleted(respone.NewCartUpdateResponse(cart))
}
