package major

import (
	"context"

	"github.com/labstack/echo/v4"
	majorEntity "go-management-auth-school/entity/major"
)

type MajorService interface {
	FindAll(ctx context.Context, params *MajorParams) (data []majorEntity.Major, err error)
	SelectAll(ctx context.Context, parameter *MajorParams) (data []majorEntity.Major, err error)
	FindOne(ctx context.Context, params *MajorParams) (data majorEntity.Major, err error)
}

type majorController struct {
	majorService MajorService
}

func NewMajorController(majorServices MajorService) majorController {
	return majorController{
		majorService: majorServices,
	}
}

func (ctrl majorController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}
