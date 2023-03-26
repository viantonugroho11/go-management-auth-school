package student

import (
	"log"
	"context"

	// "go-management-auth-school/response"

	studentEntity "go-management-auth-school/entity/student"

	"github.com/labstack/echo/v4"
)


type StudentService interface {

	SelectAll(ctx context.Context, parameter *StudentParams) (data []studentEntity.Student, err error)
	
}

type studentController struct {
	studentServices StudentService
}

func NewStudentController(studentServices StudentService) studentController {
	return studentController{
		studentServices: studentServices,
	}
}

func (ctrl studentController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("/student", ctrl.GetStudent())
}

func (ctrl studentController) GetStudent() echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		// parameter := new(StudentParams)
		// if err := c.Bind(parameter); err != nil {
		// 	return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		// }
		// if err := c.Validate(parameter); err != nil {
		// 	return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		// }
		// parameter.Offset, parameter.Limit, parameter.Page, parameter.OrderBy, parameter.Sort =
		// 	services.SetPaginationParameter(parameter.Page, parameter.Limit, studentEntity.MapOrderBy[parameter.OrderBy], parameter.Sort, studentEntity.OrderBy, studentEntity.OrderByString)
		// data, err := ctrl.studentServices.SelectAll(ctx, parameter)
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		// }
		// return c.JSON(http.StatusOK, response.SuccessResponse(data))
		log.Println("test")
		return nil
	}
		
} 