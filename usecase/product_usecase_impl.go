package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/model/entity"
	"github.com/dihanto/go-mastermind/model/web/request"
	"github.com/dihanto/go-mastermind/model/web/response"
	"github.com/dihanto/go-mastermind/repository"
	"github.com/go-playground/validator/v10"
)

type ProductUsecaseImpl struct {
	Repository repository.ProductRepository
	Db         *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewProductUsecaseImpl(repository repository.ProductRepository, db *sql.DB, validate *validator.Validate, timeout int) ProductUsecase {
	return &ProductUsecaseImpl{
		Repository: repository,
		Db:         db,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *ProductUsecaseImpl) AddProduct(ctx context.Context, request request.AddProduct) (product response.AddProduct, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	requestRepo := entity.Product{
		IdSeller:  request.IdSeller,
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		CreatedAt: int32(time.Now().Unix()),
	}
	response, err := usecase.Repository.AddProduct(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	product = helper.ToResponseAddProduct(response)

	return
}

func (usecase *ProductUsecaseImpl) GetProduct(ctx context.Context) (products []response.GetProduct, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	responses, err := usecase.Repository.GetProduct(ctx, tx)
	if err != nil {
		return
	}

	for _, product := range responses {
		response := response.GetProduct{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		products = append(products, response)
	}
	return
}

func (usecase *ProductUsecaseImpl) FindById(ctx context.Context, id int) (product response.FindById, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	response, err := usecase.Repository.FindById(ctx, tx, id)
	if err != nil {
		return
	}
	product = helper.ToResponseFindById(response)
	return
}
func (usecase *ProductUsecaseImpl) UpdateProduct(ctx context.Context, request request.UpdateProduct) (product response.UpdateProduct, err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	requestRepo := entity.Product{
		Id:        request.Id,
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		UpdatedAt: int32(time.Now().Unix()),
	}

	response, err := usecase.Repository.UpdateProduct(ctx, tx, requestRepo)
	if err != nil {
		return
	}

	product = helper.ToResponseUpdateProduct(response)

	return
}

func (usecase *ProductUsecaseImpl) DeleteProduct(ctx context.Context, id int) (err error) {
	tx, err := usecase.Db.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	deleteTime := int32(time.Now().Unix())
	err = usecase.Repository.DeleteProduct(ctx, tx, deleteTime, id)
	if err != nil {
		return
	}
	return
}
