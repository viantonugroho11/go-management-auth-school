package class

import (
	"context"
	classController "go-management-auth-school/controller/class"
	classEntity "go-management-auth-school/entity/class"
)

type ClassRepo interface {
	FindAll(ctx context.Context, params *classController.ClassParams) (data []classEntity.Class, err error)
	SelectAll(ctx context.Context, parameter *classController.ClassParams) (data []classEntity.Class, err error)
	FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error)
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
	return
}

func (service classService) FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error) {
	return
}