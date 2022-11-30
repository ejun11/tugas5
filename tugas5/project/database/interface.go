package database

import "tugas5/project/model"

type Userface interface {
	Create(user *model.User) error
	FindAll() (*[]model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(email string, user *model.User) error
	Delete(id string) error
}

type Productface interface {
	GetProducts() (*[]model.Product, error)
	CreateProduct(m *model.Product) error
	UpdateProduct(id string, m *model.Product) error
	DeleteProduct(id string, m *model.Product) error
	GetProductById(productId string) (*model.Product, error)
}

type Categoryface interface {
	GetCategories() (*[]model.Category, error)
	CreateCategory(m *model.Category) error
	UpdateCategory(id string, m *model.Category) error
	DeleteCategory(id string, m *model.Category) error
	GetCategoryById(categoryId string) (*model.Category, error)
}

type Cartface interface {
	GetCarts() (*[]model.Cart, error)
	GetUserCarts(userId string) (*[]model.Cart, error)
	GetCartById(cartId string) (*model.Cart, error)
	GetCartByProductId(userId string, productId string) (*model.Cart, error)
	CreateCart(m *model.Cart) error
	UpdateCart(cartId string, m *model.Cart) error
	DeleteCart(cartId string, m *model.Cart) error
}

type Orderface interface {
	GetOrders() (*[]model.Order, error)
	GetUserOrders(userId string) (*[]model.Order, error)
	CreateOrder(m *model.Order) error
	UpdateOrder(id string, m *model.Order) error
	DeleteOrder(id string, m *model.Order) error
	GetOrderById(orderId string) (*model.Order, error)
}

type OrderItemface interface {
	GetAllOrderItems() (*[]model.OrderItem, error)
	GetOrderItems(orderId string) (*[]model.OrderItem, error)
	CreateOrderItem(m *model.OrderItem) error
	UpdateOrderItem(id string, m *model.OrderItem) error
	DeleteOrderItem(id string, m *model.OrderItem) error
	GetOrderItemById(orderItemId string) (*model.OrderItem, error)
}

type Paymentface interface {
	GetPayments() (*[]model.Payment, error)
	GetOrderPayments(orderId string) (*[]model.Payment, error)
	GetSuccessOrderPayment(orderId string) (*model.Payment, error)
	GetPaymentById(deliveryId string) (*model.Delivery, error)
	UpdatePayment(id string, m *model.Payment) error
}

type Deliveryface interface {
	GetDeliveries() (*[]model.Delivery, error)
	GetDelivery(orderId string) (*model.Delivery, error)
	CreateDelivery(m *model.Delivery) error
	UpdateDelivery(id string, m *model.Delivery) error
	DeleteDelivery(id string, m *model.Delivery) error
	GetOrderItemById(deliveryId string) (*model.Delivery, error)
}
