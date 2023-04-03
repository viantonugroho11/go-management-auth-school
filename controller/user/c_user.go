package user

import (
	"context"
	"log"
	"net/http"

	studentEntity "go-management-auth-school/entity/student"
	userEntity "go-management-auth-school/entity/user"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)



type UserService interface {
	SelectAll(ctx context.Context, parameter *UserParams) (data []studentEntity.Student, err error)
	FindOne(ctx context.Context, parameter *UserParams) (data userEntity.User, err error)
	Create(ctx context.Context, parameter *userEntity.User) (data userEntity.User, err error)
}

type userController struct {
	userServices UserService
}

func NewUserController(service UserService) *userController {
	return &userController{
		userServices: service,
	}
}

func (ctrl userController) GetUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		
		params := new(UserParams)
		log.Println(params)
		data, err := ctrl.userServices.SelectAll(ctx, params)
		if err != nil {
			return err
		}
		return response.RespondSuccess(c,http.StatusAlreadyReported, data,nil)
	}
}