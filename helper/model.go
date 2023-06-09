package helper

import (
	"time"

	"github.com/dihanto/go-mastermind/model/entity"
	"github.com/dihanto/go-mastermind/model/web/response"
)

func ToResponseCustomerRegister(customer entity.Customer) response.CustomerRegister {
	return response.CustomerRegister{
		Email:        customer.Email,
		Name:         customer.Name,
		RegisteredAt: time.Unix(int64(customer.RegisteredAt), 0),
	}
}
func ToResponseCustomerUpdate(customer entity.Customer) response.CustomerUpdate {
	return response.CustomerUpdate{
		Email:        customer.Email,
		Name:         customer.Name,
		RegisteredAt: time.Unix(int64(customer.RegisteredAt), 0),
		UpdatedAt:    time.Unix(int64(customer.UpdatedAt), 0),
	}
}
func ToResponseSellerRegister(seller entity.Seller) response.SellerRegister {
	return response.SellerRegister{
		Email:        seller.Email,
		Name:         seller.Name,
		RegisteredAt: time.Unix(int64(seller.RegisteredAt), 0),
	}
}
func ToResponseSellerUpdate(seller entity.Seller) response.SellerUpdate {
	return response.SellerUpdate{
		Email:        seller.Email,
		Name:         seller.Name,
		RegisteredAt: time.Unix(int64(seller.RegisteredAt), 0),
		UpdatedAt:    time.Unix(int64(seller.UpdatedAt), 0),
	}
}

func ToResponseAddProduct(product entity.Product) response.AddProduct {
	return response.AddProduct{
		Id:        product.Id,
		IdSeller:  product.IdSeller,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
	}
}

func ToResponseFindById(product entity.Product) response.FindById {
	return response.FindById{
		Id:        product.Id,
		IdSeller:  product.IdSeller,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
		UpdatedAt: time.Unix(int64(product.UpdatedAt), 0),
	}
}

func ToResponseUpdateProduct(product entity.Product) response.UpdateProduct {
	return response.UpdateProduct{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
		UpdatedAt: time.Unix(int64(product.UpdatedAt), 0),
	}
}
