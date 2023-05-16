package person

import (
    "gorm.io/gorm"
)

type Appearence struct {
	gorm.Model
	PersonID 	int
	Item		string		`json:"name"`
	Tag			float32		`json:"tag"`
}