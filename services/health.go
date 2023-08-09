package services

type HealthService struct {}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) GetHealthStatus() bool {
	return true;
}