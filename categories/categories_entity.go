package categories

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}