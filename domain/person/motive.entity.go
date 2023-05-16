package person

import (
    "gorm.io/gorm"
)

type Motive struct {
	gorm.Model
	PersonID 	int
	Energy		int		`json:"energy"`
	Hunger		int		`json:"hunger"`
	Bladder		int		`json:"bladder"`
	Social		int		`json:"social"`
	Hygiene		int		`json:"hygiene"`
	Fun			int		`json:"fun"`
}