package seed

import "gorm.io/gorm"

type Seed struct {
	Name string `json:"name"`
	Run  func(*gorm.DB) error
}
