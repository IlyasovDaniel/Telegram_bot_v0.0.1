package product

type Service struct {
}

func (s *Service) Get() []Product {
	panic("unimplemented")
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) List() []Product {
	return allProduct
}
