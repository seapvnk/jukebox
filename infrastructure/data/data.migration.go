package data

import (
	db "github.com/seapvnk/jukebox/infrastructure/database"
	place "github.com/seapvnk/jukebox/domain/place"	
)

func Migration() {
	db.DB.AutoMigrate(&place.Place{})
}

