package parent

import ()

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