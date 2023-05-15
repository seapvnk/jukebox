package place

import (
    "gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Map			string	`json:"map"`
	Width		int		`json:"width"`
	Height		int		`json:"height"`
}