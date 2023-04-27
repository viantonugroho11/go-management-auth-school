package mapping_course

import (
	"context"

	"github.com/labstack/echo/v4"

	mapCourseEntity "go-management-auth-school/entity/mapping_course"
)

type MappingCourseService interface {
	FindAll(ctx context.Context, params *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	SelectAll(ctx context.Context, parameter *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	FindOne(ctx context.Context, params *MappingCourseParams) (data mapCourseEntity.MappingCourse, err error)
}

type mappingCourseController struct {
	mappingCourseService MappingCourseService
}

func NewMappingCourseController(mappingCourseService MappingCourseService) mappingCourseController {
	return mappingCourseController{
		mappingCourseService: mappingCourseService,
	}
}

func (ctrl mappingCourseController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}
