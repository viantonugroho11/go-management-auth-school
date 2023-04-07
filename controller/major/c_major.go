package major

import "github.com/labstack/echo/v4"


type MajorService interface {
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