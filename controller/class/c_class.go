package class

import (
	"context"

	"github.com/labstack/echo/v4"

	classEntity "go-management-auth-school/entity/class"
)

type ClassService interface {
	FindAll(ctx context.Context, params *ClassParams) (data []classEntity.Class, err error)
	SelectAll(ctx context.Context, parameter *ClassParams) (data []classEntity.Class, err error)
	FindOne(ctx context.Context, params *ClassParams) (data classEntity.Class, err error)
}

type classController struct {
	classService ClassService
}

func NewClassController(classService ClassService) classController {
	return classController{
		classService: classService,
	}
}

func (ctrl classController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}