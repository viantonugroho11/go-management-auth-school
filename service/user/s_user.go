package user

import (
	"context"

	userRequset "go-management-auth-school/controller/user"
	userEntity "go-management-auth-school/entity/user"

	helperStr "go-management-auth-school/helper/str"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	SelectAll(ctx context.Context, parameter *userRequset.UserParams) (data []userEntity.User, err error)
	FindOne(ctx context.Context, parameter *userRequset.UserParams) (data userEntity.User, err error)
	Create(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (res string,err error)
	UpdateUsername(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error)
	UpdatePassword(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
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

func (service userService) FindOne(ctx context.Context, parameter *userRequset.UserParams) (data userEntity.User, err error) {
	data, err = service.userRepo.FindOne(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}

func (service userService) Create(ctx context.Context, input *userEntity.User) (data userEntity.User, err error) {

	if !helperStr.ValidatePassword(input.Password) {
		return
	}
	// check if username already exist
	data, err = service.userRepo.FindOne(ctx, &userRequset.UserParams{
		Username: input.Username,
	})
	if err != nil {
		return
	}
	if data.ID != "" {
		return
	}


	tx , err := service.userRepo.CreateTx(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback()
	// permission 0 = user
	// permission 1 = guru
	// permission 2 = admin
	data.IdentityID, err = service.userRepo.Create(ctx, tx, input)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}

	tx.Commit()
	return
}

// update
func (service userService) UpdateUsername(ctx context.Context, input *userEntity.User) (err error) {
	// check if username already exist
	data, err := service.userRepo.FindOne(ctx, &userRequset.UserParams{
		IdentityID: input.IdentityID,
	})
	if err != nil {
		return
	}

	checkData,err := service.userRepo.FindOne(ctx, &userRequset.UserParams{
		Username: input.Username,
	})

	if err != nil {
		return
	}
	if checkData.IdentityID != input.IdentityID && checkData.Username == data.Username {
		return
	}

	tx , err := service.userRepo.CreateTx(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback()
	err = service.userRepo.UpdateUsername(ctx, tx, &userEntity.User{
		IdentityID: input.IdentityID,
		Username: input.Username,
	})
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}

	tx.Commit()


	
	
	return
}
