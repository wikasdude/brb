package usecase

import (
	"brb-midsvc-platform/internal/domain"
	"brb-midsvc-platform/internal/repository"
	"errors"
	"strconv"
)

type VendorUsecase interface {
	CreateVendor(vendor *domain.Vendor) error
	UpdateVendor(vendor *domain.Vendor) error
	GetVendorByID(id int64) (*domain.Vendor, error)
	ListVendors() ([]domain.Vendor, error)
}

type vendorUsecase struct {
	repo repository.VendorRepository
}

func NewVendorUsecase(repo repository.VendorRepository) VendorUsecase {
	return &vendorUsecase{repo: repo}
}

func (u *vendorUsecase) CreateVendor(vendor *domain.Vendor) error {
	if vendor.Name == "" {
		return errors.New("vendor name is required")
	}
	return u.repo.Create(vendor)
}

func (u *vendorUsecase) UpdateVendor(vendor *domain.Vendor) error {
	existing, err := u.repo.GetByID(vendor.ID)
	if err != nil {
		return err
	}
	existing.Name = vendor.Name
	//existing.Description = vendor.Description
	return u.repo.Update(existing)
}

func (u *vendorUsecase) GetVendorByID(id int64) (*domain.Vendor, error) {
	return u.repo.GetByID(strconv.Itoa(int(id)))
}

func (u *vendorUsecase) ListVendors() ([]domain.Vendor, error) {
	return u.repo.ListAll()
}
