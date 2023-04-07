package admin

import "github.com/labstack/echo/v4"

type AdminService interface {
}

type adminController struct {
	adminService AdminService
}

func NewAdminController(adminService AdminService) adminController {
	return adminController{
		adminService: adminService,
	}
}

func (ctrl adminController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}


