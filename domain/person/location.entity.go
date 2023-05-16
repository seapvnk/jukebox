package person

import (
    "gorm.io/gorm"
)

type Location struct {
	gorm.Model
	PersonID 	int
	Map			string		`json:"map"`
	X			float32		`json:"x"`
	Y			float32		`json:"y"`
}