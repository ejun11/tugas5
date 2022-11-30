package gorm

import (
	"time"
	database "tugas5/project/database"
	"tugas5/project/model"

	"gorm.io/gorm"
)

type userface struct {
	db *gorm.DB
}

type cartface struct {
	db *gorm.DB
}

type categoryface struct {
	db *gorm.DB
}

type productface struct {
	db *gorm.DB
}

func NewCartdb(db *gorm.DB) database.Cartface {
	return &cartface{
		db: db,
	}
}

func NewUserface(db *gorm.DB) database.Userface {
	return &userface{
		db: db,
	}
}

func (u *userface) Create(user *model.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *userface) FindAll() (*[]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userface) FindById(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userface) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email= ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *userface) Update(email string, user *model.User) error {
	return p.db.Where("email= ?", email).Updates(user).Error
}

func (u *userface) Delete(id string) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (p *cartface) GetCarts() (*[]model.Cart, error) {
	var carts []model.Cart

	err := p.db.Where("deleted NULL").Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return &carts, nil
}

func (p *cartface) GetUserCarts(userId string) (*[]model.Cart, error) {
	var carts []model.Cart

	err := p.db.Where("user_id = ?", userId).Where("deleted NULL").Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return &carts, nil
}

func (p *cartface) GetCartById(cartId string) (*model.Cart, error) {
	var cart model.Cart
	err := p.db.Joins("User").Joins("Product").Where("carts.id = ?", cartId).Where("carts.deleted NULL").First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (p *cartface) GetCartByProductId(userId string, productId string) (*model.Cart, error) {
	var cart model.Cart
	err := p.db.Where("carts.user_id = ?", userId).Where("carts.product_id = ?", productId).Where("carts.deleted NULL").First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (p *cartface) CreateCart(cart *model.Cart) error {
	return p.db.Create(cart).Error
}

func (p *cartface) UpdateCart(id string, cart *model.Cart) error {
	return p.db.Where("id = ?", id).Updates(cart).Error
}

func (p *cartface) DeleteCart(id string, cart *model.Cart) error {
	return p.db.Model(&cart).Where("id = ?", id).Update("deleted", time.Now()).Error
}

func NewCategoryface(db *gorm.DB) database.Categoryface {
	return &categoryface{
		db: db,
	}
}

func (c *categoryface) GetCategories() (*[]model.Category, error) {
	var categories []model.Category

	err := c.db.Where("deleted NULL").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (c *categoryface) CreateCategory(category *model.Category) error {
	return c.db.Create(category).Error
}

func (p *categoryface) UpdateCategory(id string, category *model.Category) error {
	return p.db.Where("id = ?", id).Updates(category).Error
}

func (p *categoryface) DeleteCategory(id string, category *model.Category) error {
	return p.db.Model(&category).Where("id = ?", id).Update("deleted", time.Now()).Error
}

func (c *categoryface) GetCategoryById(categoryId string) (*model.Category, error) {
	var category model.Category
	err := c.db.Where("categories.id=?", categoryId).Where("categories.deleted NULL").First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func NewProductface(db *gorm.DB) database.Productface {
	return &productface{
		db: db,
	}
}

func (p *productface) GetProducts() (*[]model.Product, error) {
	var products []model.Product

	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (p *productface) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productface) UpdateProduct(id string, product *model.Product) error {
	return p.db.Where("id = ?", id).Updates(product).Error
}

func (p *productface) DeleteProduct(id string, product *model.Product) error {
	return p.db.Model(&product).Where("id = ?", id).Update("deleted", time.Now()).Error
}

func (p *productface) GetProductById(productId string) (*model.Product, error) {
	var product model.Product
	err := p.db.Joins("Category").Where("products.id=?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
