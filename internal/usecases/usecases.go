package usecases

type Usecases interface {
	UserUsecases
}

type UsecasesService struct {
	Repo Repository
}
