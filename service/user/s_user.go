package user

import (
	"context"

	userEntity "go-management-auth-school/entity/user"
	userRequset "go-management-auth-school/controller/user"
)

type UserRepo interface {
	SelectAll(ctx context.Context, parameter *userRequset.UserParams) (data []userEntity.User, err error)
	FindOne(ctx context.Context, parameter *userRequset.UserParams) (data userEntity.User, err error)
}


type userService struct {
	userRepo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{
		userRepo: repo,
	}
}

func (service userService)SelectAll(ctx context.Context, parameter *userRequset.UserParams) (data []userEntity.User, err error) {
	// parameter.Offset, parameter.Limit, parameter.Page, parameter.OrderBy, parameter.Sort =
	// 	service.SetPaginationParameter(parameter.Page, parameter.Limit, studentEntity.MapOrderBy[parameter.OrderBy], parameter.Sort, studentEntity.OrderBy, studentEntity.OrderByString)

	data, err = service.userRepo.SelectAll(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}

func (service userService)FindOne(ctx context.Context, parameter *userRequset.UserParams) (data userEntity.User, err error) {
	data, err = service.userRepo.FindOne(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}
