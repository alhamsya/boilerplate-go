package omdb

func (o *OMDB) GetListMovie(page int, search string) (*OMDBList, error) {
	return &OMDBList{}, nil
}

func (o *OMDB) GetDetailMovie(omdbID string) (*OMDBDetail, error) {
	return &OMDBDetail{}, nil
}
