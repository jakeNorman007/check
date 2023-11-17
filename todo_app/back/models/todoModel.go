package models

import "gorm.io/gorm"

// gorm.Model automatically adds fields for us *make better note*
type Todo struct {
    gorm.Model
    Body    string
}
