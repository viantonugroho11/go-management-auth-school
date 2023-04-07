package lesson

import "github.com/labstack/echo/v4"

type LessonService interface {
}

type lessonController struct {
	lessonService LessonService
}

func NewLessonController(lessonService LessonService) lessonController {
	return lessonController{
		lessonService: lessonService,
	}
}

func (ctrl lessonController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}