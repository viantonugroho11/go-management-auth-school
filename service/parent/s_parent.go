package parent

import (
	"context"


	parentController "go-management-auth-school/controller/parent"
	parentEntity "go-management-auth-school/entity/parent"

)

type ParentRepo interface {
	FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error)
	SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error)
	FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error)
}


type parentService struct {
	parentRepo ParentRepo
}

func NewParentService(repo ParentRepo) *parentService {
	return &parentService{
		parentRepo: repo,
	}
}

func (service parentService) FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	return
}

func (service parentService) SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	return
}

func (service parentService) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	return
}