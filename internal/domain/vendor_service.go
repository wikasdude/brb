package domain

type VendorService struct {
	ID        string `gorm:"primaryKey"`
	VendorID  string
	ServiceID string
	Active    bool
}
