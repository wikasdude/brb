package usecase

import (
	"brb-midsvc-platform/internal/domain"
	"brb-midsvc-platform/internal/repository"
	"errors"
)

type ServiceUsecase interface {
	CreateService(service *domain.Service) error
	UpdateService(service *domain.Service) error
	GetServiceByID(id int64) (*domain.Service, error)
	ListServices() ([]domain.Service, error)
	ToggleServiceAvailability(id int64, active bool) error
}

type serviceUsecase struct {
	repo repository.ServiceRepository
}

func NewServiceUsecase(repo repository.ServiceRepository) ServiceUsecase {
	return &serviceUsecase{repo: repo}
}

func (u *serviceUsecase) CreateService(service *domain.Service) error {
	if service.Name == "" {
		return errors.New("service name is required")
	}
	return u.repo.Create(service)
}

func (u *serviceUsecase) UpdateService(service *domain.Service) error {
	existing, err := u.repo.GetByID(service.ID)
	if err != nil {
		return err
	}
	existing.Name = service.Name
	existing.Description = service.Description
	existing.Active = service.Active
	return u.repo.Update(existing)
}

func (u *serviceUsecase) GetServiceByID(id int64) (*domain.Service, error) {
	return u.repo.GetByID(id)
}

func (u *serviceUsecase) ListServices() ([]domain.Service, error) {
	return u.repo.ListAll()
}

func (u *serviceUsecase) ToggleServiceAvailability(id int64, active bool) error {
	service, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	service.Active = active
	return u.repo.Update(service)
}
