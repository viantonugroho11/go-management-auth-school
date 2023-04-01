package admin

import ()

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


