package data

import (
	db "github.com/seapvnk/jukebox/infrastructure/database"
	place "github.com/seapvnk/jukebox/domain/place"	
	person "github.com/seapvnk/jukebox/domain/person"	
	mood "github.com/seapvnk/jukebox/domain/mood"	
)

func Migration() {
	db.DB.AutoMigrate(&place.Place{})
	db.DB.AutoMigrate(&mood.Mood{})
	db.DB.AutoMigrate(&person.Appearence{})
	db.DB.AutoMigrate(&person.Person{})
	db.DB.AutoMigrate(&person.Location{})
	db.DB.AutoMigrate(&person.Motive{})
}

