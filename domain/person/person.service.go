package person

func GetPersonService(id int) []Person {
	return PersonGetOne(id)
}

func GetAllPeopleService() []Person {
	return PersonGet()
}

func CreatePersonService(dto CreatePersonDto) []Person {
	people := []Person{ dto.Person() }
	PeopleCreate(&people)

	return people
}