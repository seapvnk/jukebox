package place

type CreatePlaceDto struct {
	Name		string	`form:"name" binding:"required"`
	Description	string	`form:"description" binding:"required"`
	Map			string	`form:"map" binding:"required"`
	Width		int		`form:"width" binding:"required"`
	Height		int		`form:"height" binding:"required"`
}

func (dto *CreatePlaceDto) Place() Place {
	var place Place

	place.Name = dto.Name
	place.Description = dto.Description
	place.Map = dto.Map
	place.Width = dto.Width
	place.Height = dto.Height

	return place
}

type GetPlaceDto struct {
	ID	int	`uri:"id" binding:"required"`
}