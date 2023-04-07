package admin

type AdminRepo interface {
}


type adminService struct {
	adminRepo AdminRepo
}

func NewAdminService(repo AdminRepo) *adminService {
	return &adminService{
		adminRepo: repo,
	}
}