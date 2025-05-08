package domain

type Booking struct {
	ID        int    `gorm:"primaryKey"`
	UserID    string `json:"user_id"`
	VendorID  int64  `json:"vendor_id"`
	ServiceID string `json:"service_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"` // pending, confirmed, completed
}
