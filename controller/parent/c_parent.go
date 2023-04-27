package parent

import (
	"context"

	"github.com/labstack/echo/v4"

	parentEntity "go-management-auth-school/entity/parent"
)

type ParentService interface {
	FindAll(ctx context.Context, params *ParentParams) (data []parentEntity.Parent, err error)
	SelectAll(ctx context.Context, parameter *ParentParams) (data []parentEntity.Parent, err error)
	FindOne(ctx context.Context, params *ParentParams) (data parentEntity.Parent, err error)
}

type parentController struct {
	parentService ParentService
}

func NewParentController(parentService ParentService) parentController {
	return parentController{
		parentService: parentService,
	}
}

func (ctrl parentController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}