package omdb

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type OMDB struct {
	Cfg *config.ServiceConfig
}

type OMDBList struct {
	Search       []Search `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
	Error        string   `json:"Error"`
}
type Search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type OMDBDetail struct {
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	Ratings    []Ratings `json:"Ratings"`
	Metascore  string    `json:"Metascore"`
	ImdbRating string    `json:"imdbRating"`
	ImdbVotes  string    `json:"imdbVotes"`
	ImdbID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	DVD        string    `json:"DVD"`
	BoxOffice  string    `json:"BoxOffice"`
	Production string    `json:"Production"`
	Website    string    `json:"Website"`
	Response   string    `json:"Response"`
	Error      string    `json:"Error"`
}
type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
