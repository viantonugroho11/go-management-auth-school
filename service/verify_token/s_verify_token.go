package verify_token

import (
	"context"
	verifyTokenRequest "go-management-auth-school/controller/verify_token"
	verifyTokenEntity "go-management-auth-school/entity/verify_token"

	"github.com/jmoiron/sqlx"
)


type VerifyTokenRepo interface {
	FindOne(ctx context.Context, parameter *verifyTokenRequest.VerifyTokenParams)(data verifyTokenEntity.VerifyToken, err error)
	Create(ctx context.Context, tx *sqlx.Tx, parameter *verifyTokenEntity.VerifyToken) (err error)
	CreateTx(ctx context.Context)(tx *sqlx.Tx, err error)
	Delete(ctx context.Context, tx *sqlx.Tx, parameter *verifyTokenEntity.VerifyToken)(err error)
}

type verifyTokenService struct {
	verifyTokenRepo VerifyTokenRepo
}

func NewVerifyTokenService(repo VerifyTokenRepo) *verifyTokenService {
	return &verifyTokenService{
		verifyTokenRepo: repo,
	}
}

func (u *verifyTokenService) FindOne(ctx context.Context, parameter *verifyTokenRequest.VerifyTokenParams) (data verifyTokenEntity.VerifyToken, err error) {
	data, err = u.verifyTokenRepo.FindOne(ctx, parameter)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (u *verifyTokenService) Create(ctx context.Context, parameter *verifyTokenEntity.VerifyToken) (err error) {
	tx, err := u.verifyTokenRepo.CreateTx(ctx)
	if err != nil {
		return err
	}

	err = u.verifyTokenRepo.Create(ctx, tx, parameter)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *verifyTokenService) Delete(ctx context.Context, parameter *verifyTokenEntity.VerifyToken) (err error) {
	tx , err := u.verifyTokenRepo.CreateTx(ctx)
	if err != nil {
		return err
	}

	err = u.verifyTokenRepo.Delete(ctx, tx, parameter)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
