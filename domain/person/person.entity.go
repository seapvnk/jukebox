package person

import (
    "gorm.io/gorm"
	mood "github.com/seapvnk/jukebox/domain/mood"	
)

type Person struct {
	gorm.Model
	Appearence		[]Appearence	`json:"appearence" gorm:"foreignKey:PersonID"`
	Moods			[]mood.Mood		`json:"moods" gorm:"many2many:person_mood;"`
	Motive			Motive			`json:"motive" gorm:"foreignKey:PersonID"`
	Location		Location		`json:"location" gorm:"foreignKey:PersonID"`
	ProfilePic		string			`json:"profilePic"`
	FirstName		string			`json:"firstName"`
	LastName		string			`json:"lastName"`
	Bio				string			`json:"bio"`
	Money			float32			`json:"money"`

}