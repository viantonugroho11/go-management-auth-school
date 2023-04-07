package class

import "github.com/labstack/echo/v4"

type ClassService interface {
	// FindOne(id int)
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