package models

import "gorm.io/gorm"

type ObjectModel struct {
	gorm.Model
	Name string `gorm:unique`
	Value string
}
