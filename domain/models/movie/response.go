package modelMovie

type RespListMovie struct {
	Items []Items `json:"items"`
	Total int64     `json:"total"`
}

type Items struct {
	Title   string `json:"title"`
	Year    string `json:"year"`
	MovieID string `json:"movie_id"`
	Types   string `json:"types"`
	Poster  string `json:"poster"`
}
