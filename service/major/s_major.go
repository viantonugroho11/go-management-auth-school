package major

type MajorRepo interface {
}


type majorService struct {
	majorRepo MajorRepo
}

func NewMajorService(repo MajorRepo) *majorService {
	return &majorService{
		majorRepo: repo,
	}
}
