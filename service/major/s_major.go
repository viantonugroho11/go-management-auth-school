package major

import (
	"context"
	majorController "go-management-auth-school/controller/major"
	majorEntity "go-management-auth-school/entity/major"
)

type MajorRepo interface {
	FindAll(ctx context.Context, params *majorController.MajorParams) (data []majorEntity.Major, err error)
	SelectAll(ctx context.Context, parameter *majorController.MajorParams) (data []majorEntity.Major, err error)
	FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error)
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
	return
}

func (service majorService) FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error) {
	return
}
