package person

import (
    "gorm.io/gorm"
)

type Mood struct {
	gorm.Model
	Color		int		`json:"color"`
	Name		int		`json:"name"`
}