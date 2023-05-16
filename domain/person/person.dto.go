package person

type CreatePersonDto struct {
	Appearence		[]Appearence	`form:"appearence" binding:"required`
	FirstName		string			`form:"firstName" binding:"required`
	LastName		string			`form:"lastName" binding:"required`
	Bio				string			`form:"bio" binding:"required`
	ProfilePic		string			`form:"profilePic" binding:"required`
	Money			float32			`form:"money" binding:"required`
}

func (dto *CreatePersonDto) Person() Person {
	var person Person

	person.FirstName = dto.FirstName
	person.LastName = dto.LastName
	person.Money = dto.Money
	person.Bio = dto.Bio
	person.ProfilePic = dto.ProfilePic
	person.Appearence = dto.Appearence

	
	defaultMotiveValue := 100
	person.Motive.Energy = defaultMotiveValue
	person.Motive.Hunger = defaultMotiveValue
	person.Motive.Bladder = defaultMotiveValue
	person.Motive.Social = defaultMotiveValue
	person.Motive.Hygiene = defaultMotiveValue
	person.Motive.Fun = defaultMotiveValue

	return person
}

type GetPersonDto struct {
	ID	int	`uri:"id" binding:"required"`
}