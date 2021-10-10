package modelMovie

type RespListMovie struct {
	Items []Items `json:"items"`
	Total int64   `json:"total"`
}

type RespDetailMovie struct {
	Title      string    `json:"title"`
	Year       string    `json:"year"`
	Rated      string    `json:"rated"`
	Released   string    `json:"released"`
	Runtime    string    `json:"runtime"`
	Genre      string    `json:"genre"`
	Director   string    `json:"director"`
	Writer     string    `json:"writer"`
	Actors     string    `json:"actors"`
	Plot       string    `json:"plot"`
	Language   string    `json:"language"`
	Country    string    `json:"country"`
	Awards     string    `json:"awards"`
	Poster     string    `json:"poster"`
	Ratings    []Ratings `json:"ratings"`
	MetaScore  string    `json:"meta_score"`
	ImdbRating string    `json:"imdb_rating"`
	ImdbVotes  string    `json:"imdb_votes"`
	ImdbID     string    `json:"imdb_id"`
	Type       string    `json:"type"`
	DVD        string    `json:"dvd"`
	BoxOffice  string    `json:"box_office"`
	Production string    `json:"production"`
	Website    string    `json:"website"`
}
type Ratings struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

type Items struct {
	Title   string `json:"title"`
	Year    string `json:"year"`
	MovieID string `json:"movie_id"`
	Types   string `json:"types"`
	Poster  string `json:"poster"`
}
