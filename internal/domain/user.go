package domain

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string
	Role     string // "admin" or "customer"
	Password string
}
