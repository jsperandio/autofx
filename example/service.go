package example

type Service struct {
	options Options
	db      Store
}

func NewService(opt Options, db Store) *Service {
	return &Service{
		options: opt,
		db:      db,
	}
}

func (s *Service) DoSomething(something string) string {
	result, err := s.db.Get(something)
	if err != nil {
		return ""
	}
	return s.options.SomeOption + result
}

func (s *Service) DoSomethingElse(somethingElse int) int {
	return somethingElse + 10
}
