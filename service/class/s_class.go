package class

import (
	"context"
	classController "go-management-auth-school/controller/class"
	classEntity "go-management-auth-school/entity/class"

	"github.com/jmoiron/sqlx"
)

type ClassRepo interface {
	FindAll(ctx context.Context, params *classController.ClassParams) (data []classEntity.Class, err error)
	SelectAll(ctx context.Context, parameter *classController.ClassParams) (data []classEntity.Class, err error)
	FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *classEntity.Class) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type classService struct {
	classRepo ClassRepo
}

func NewClassService(repo ClassRepo) *classService {
	return &classService{
		classRepo: repo,
	}
}

func (service classService) FindAll(ctx context.Context, params *classController.ClassParams) (data []classEntity.Class, err error) {
	return
}

func (service classService) SelectAll(ctx context.Context, parameter *classController.ClassParams) (data []classEntity.Class, err error) {
	data, err = service.classRepo.SelectAll(ctx, parameter)
	if err != nil {
		return data, err
	}
	return
}

func (service classService) FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error) {
	data, err = service.classRepo.FindOne(ctx, params)
	if err != nil {
		return data, err
	}
	return
}

func (service classService) Create(ctx context.Context, params *classEntity.Class) (err error) {
	tx, err := service.classRepo.CreateTx(ctx)
	if err != nil {
		return err
	}

	err = service.classRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return
}
