package place

import (
	db "github.com/seapvnk/jukebox/infrastructure/database"
)

func PlaceGet() []Place {
	var places []Place
	db.DB.Find(&places)

	return places
}

func PlaceGetOne(id int) []Place {
	var places []Place
	db.DB.Find(&places, id)

	return places
}

func PlacesCreate(places *[]Place) []Place {
	db.DB.CreateInBatches(&places, db.DATABASE_BATCH_SIZE)

	return *places
}