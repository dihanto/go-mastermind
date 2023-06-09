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
	"github.com/google/uuid"
)

type CustomerUsecaseImpl struct {
	Repository repository.CustomerRepository
	Database   *sql.DB
	Validate   *validator.Validate
	Timeout    int
}

func NewCustomerUsecaseImpl(repository repository.CustomerRepository, database *sql.DB, validate *validator.Validate, timeout int) CustomerUsecase {
	return &CustomerUsecaseImpl{
		Repository: repository,
		Database:   database,
		Validate:   validate,
		Timeout:    timeout,
	}
}

func (usecase *CustomerUsecaseImpl) RegisterCustomer(ctx context.Context, request request.CustomerRegister) (response response.CustomerRegister, err error) {
	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return
	}

	customer := entity.Customer{
		Id:           uuid.New(),
		Email:        request.Email,
		Name:         request.Name,
		Password:     password,
		RegisteredAt: int32(time.Now().Unix()),
	}
	customerResponse, err := usecase.Repository.RegisterCustomer(ctx, tx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerRegister(customerResponse)

	return

}

func (usecase *CustomerUsecaseImpl) LoginCustomer(ctx context.Context, email string, password string) (id uuid.UUID, result bool, err error) {
	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	id, passwordHashed, err := usecase.Repository.LoginCustomer(ctx, tx, email)
	if err != nil {
		return
	}

	result, err = helper.CheckPasswordHash(passwordHashed, password)

	if !result {
		return
	}
	return

}

func (usecase *CustomerUsecaseImpl) UpdateCustomer(ctx context.Context, request request.CustomerUpdate) (response response.CustomerUpdate, err error) {
	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	customer := entity.Customer{
		Name:      request.Name,
		Email:     request.Email,
		UpdatedAt: int32(time.Now().Unix()),
	}

	customerResponse, err := usecase.Repository.UpdateCustomer(ctx, tx, customer)
	if err != nil {
		return
	}

	response = helper.ToResponseCustomerUpdate(customerResponse)

	return
}

func (usecase *CustomerUsecaseImpl) DeleteCustomer(ctx context.Context, email string) (err error) {
	tx, err := usecase.Database.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	deletedTime := int32(time.Now().Unix())

	err = usecase.Repository.DeleteCustomer(ctx, tx, email, deletedTime)
	if err != nil {
		return
	}

	return
}
