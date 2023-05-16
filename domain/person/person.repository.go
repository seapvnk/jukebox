package person

import (
	db "github.com/seapvnk/jukebox/infrastructure/database"
)

func PersonGet() []Person {
	var people []Person
	db.DB.Find(&people)

	return people
}

func PersonGetOne(id int) []Person {
	var people []Person
	db.DB.Find(&people, id)

	return people
}

func PeopleCreate(people *[]Person) []Person {
	db.DB.Omit("Motive").
		  Omit("Location").
		  Omit("Mood").
		  CreateInBatches(&people, db.DATABASE_BATCH_SIZE)

	return *people
}