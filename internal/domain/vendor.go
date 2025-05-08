package domain

type Vendor struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
	Phone string
}
type VendorSummary struct {
	VendorID      int64
	TotalBookings int
	StatusCounts  map[string]int
}
