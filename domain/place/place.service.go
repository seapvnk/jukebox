package place

func GetPlaceService(id int) []Place {
	return PlaceGetOne(id)
}

func GetAllPlacesService() []Place {
	return PlaceGet()
}

func CreatePlaceService(dto CreatePlaceDto) []Place {
	places := []Place{ dto.Place() }
	PlacesCreate(&places)

	return places
}