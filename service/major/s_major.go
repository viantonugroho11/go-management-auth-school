package major

import (
	"context"
	majorController "go-management-auth-school/controller/major"
	majorEntity "go-management-auth-school/entity/major"

	"github.com/jmoiron/sqlx"
)

type MajorRepo interface {
	FindAll(ctx context.Context, params *majorController.MajorParams) (data []majorEntity.Major, err error)
	SelectAll(ctx context.Context, parameter *majorController.MajorParams) (data []majorEntity.Major, err error)
	FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *majorEntity.Major) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type majorService struct {
	majorRepo MajorRepo
}

func NewMajorService(repo MajorRepo) *majorService {
	return &majorService{
		majorRepo: repo,
	}
}

func (service majorService) FindAll(ctx context.Context, params *majorController.MajorParams) (data []majorEntity.Major, err error) {
	return
}

func (service majorService) SelectAll(ctx context.Context, parameter *majorController.MajorParams) (data []majorEntity.Major, err error) {
	data, err = service.majorRepo.SelectAll(ctx, parameter)
	if err != nil {
		return nil, err
	}
	return
}

func (service majorService) FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error) {
	data, err = service.majorRepo.FindOne(ctx, params)
	if err != nil {
		return data, err
	}
	return
}

func (service majorService) Create(ctx context.Context, params *majorEntity.Major) (err error) {
	tx, err := service.majorRepo.CreateTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = service.majorRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return
}

func (service majorService) Update(ctx context.Context, params *majorEntity.Major) (err error) {
	return
}
