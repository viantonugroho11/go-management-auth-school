package parent


type ParentRepo interface {
}


type parentService struct {
	parentRepo ParentRepo
}

func NewParentService(repo ParentRepo) *parentService {
	return &parentService{
		parentRepo: repo,
	}
}