package ward
type WardService interface {
}

type wardController struct {
	wardService WardService
}

func NewWardController(wardServices WardService) wardController {
	return wardController{
		wardService: wardServices,
	}
}