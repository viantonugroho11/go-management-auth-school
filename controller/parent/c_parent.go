package parent

import "github.com/labstack/echo/v4"

type ParentService interface {
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