package class

type ClassRepo interface {
}


type classService struct {
	classRepo ClassRepo
}

func NewClassService(repo ClassRepo) *classService {
	return &classService{
		classRepo: repo,
	}
}