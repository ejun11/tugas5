package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tugas5/project/params"
	respone "tugas5/project/respone"
	"tugas5/project/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc *service.AuthServices
}

type UserHandler struct {
	svc *service.UserServices
}

type CategoryHandler struct {
	svc *service.CategoryService
}

type ProductHandler struct {
	svc *service.ProductService
}

type CartHandler struct {
	svc *service.CartService
}

func NewAuthHandler(svc *service.AuthServices) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func NewUserHandler(svc *service.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}

func NewCartHandler(svc *service.CartService) *CartHandler {
	return &CartHandler{
		svc: svc,
	}
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req params.AuthRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := a.svc.Register(&req)
	fmt.Println(resp)
	WriteJsonResponseGin(c, resp)
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req params.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := a.svc.Login(&req)
	WriteJsonResponseGin(c, resp)
}

func (a *UserHandler) Profile(c *gin.Context) {
	fmt.Println("Login ", c.GetString("USER_EMAIL"))
	resp := a.svc.FindByEmail(c.GetString("USER_EMAIL"))
	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) Create(c *gin.Context) {
	var req params.UserCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := u.svc.Create(&req)
	fmt.Println(resp)
	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) FindAll(c *gin.Context) {
	fmt.Println("Login ", c.GetString("USER_EMAIL"))

	resp := u.svc.FindAll()

	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) FindByEmail(c *gin.Context) {
	fmt.Println("Logiin ", c.GetString("USER_EMAIL"))
	email, isExist := c.Params.Get("email")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("email not found"))
		return
	}

	resp := u.svc.FindByEmail(email)

	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) Update(c *gin.Context) {
	email := c.GetString("USER_EMAIL")

	var req params.UserUpdate

	fmt.Println((req))
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		user := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, user)
		return
	}

	resp := u.svc.Update(email, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CategoryHandler) GetCategories(c *gin.Context) {
	resp := m.svc.GetCategories()

	WriteJsonResponseGin(c, resp)

}

func (m *CategoryHandler) CreateCategory(c *gin.Context) {
	var req params.CategoryCreate

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.CreateCategory(&req)

	WriteJsonResponseGin(c, resp)
}

func (m *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("category not found"))
		return
	}

	var req params.CategoryUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := m.svc.UpdateCategory(categoryId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("category not found"))
		return
	}

	resp := m.svc.DeleteCategory(categoryId)

	WriteJsonResponseGin(c, resp)
}

func (p *CategoryHandler) GetCategoryById(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("categoryId not found"))
		return
	}

	resp := p.svc.GetCategoryById(categoryId)
	WriteJsonResponseGin(c, resp)
}

func (p *ProductHandler) GetProducts(c *gin.Context) {
	resp := p.svc.GetProducts()
	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) CreateProduct(c *gin.Context) {
	var req params.ProductCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.CreateProduct(&req)

	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) UpdateProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("product not found"))
		return
	}

	var req params.ProductUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := m.svc.UpdateProduct(productId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("product not found"))
		return
	}

	resp := m.svc.DeleteProduct(productId)

	WriteJsonResponseGin(c, resp)
}

func (p *ProductHandler) GetProductById(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("productId not found"))
		return
	}

	resp := p.svc.GetProductById(productId)
	WriteJsonResponseGin(c, resp)
}

func (p *CartHandler) GetCarts(c *gin.Context) {
	resp := p.svc.GetCarts()

	WriteJsonResponseGin(c, resp)
}

func (p *CartHandler) GetUserCarts(c *gin.Context) {
	userId := c.GetString("USER_ID")
	resp := p.svc.GetUserCarts(userId)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) CreateCart(c *gin.Context) {
	var req params.CartCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.CreateCart(userId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) UpdateCart(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("cart not found"))
		return
	}

	var req params.CartUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, respone.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := respone.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.UpdateCart(userId, cartId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) DeleteCart(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("cart not found"))
		return
	}

	resp := m.svc.DeleteCart(cartId)

	WriteJsonResponseGin(c, resp)
}

func (p *CartHandler) GetCartById(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, respone.ErrBadRequest("cart id not found"))
		return
	}

	resp := p.svc.GetCartById(cartId)
	WriteJsonResponseGin(c, resp)
}

func WriteJsonResponse(w http.ResponseWriter, payload *respone.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	json.NewEncoder(w).Encode(payload)
}

func WriteJsonResponseGin(c *gin.Context, payload *respone.Response) {
	c.JSON(payload.Status, payload)
}

func WriteErrorJsonResponseGin(c *gin.Context, payload *respone.Response) {
	c.AbortWithStatusJSON(payload.Status, payload)
}
