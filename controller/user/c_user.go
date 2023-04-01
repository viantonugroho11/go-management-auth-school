package user

import (
	"context"
	"log"
	"net/http"

	studentEntity "go-management-auth-school/entity/student"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)



type userService interface {
	SelectAll(ctx context.Context, parameter *UserParams) (data []studentEntity.Student, err error)
	FindOne(ctx context.Context, parameter *UserParams) (data studentEntity.Student, err error)
}

type userController struct {
	userServices userService
}

func NewUserController(service userService) *userController {
	return &userController{
		userServices: service,
	}
}

func (ctrl userController) GetStudent() echo.HandlerFunc{
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