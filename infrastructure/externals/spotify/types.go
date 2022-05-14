package spotify

import "github.com/alhamsya/boilerplate-go/lib/managers/config"

type Spotify struct {
	Cfg *config.ServiceConfig
}

type Profile struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	ExplicitContent ExplicitContent `json:"explicit_content"`
	ExternalUrls    ExternalUrls    `json:"external_urls"`
	Followers       Followers       `json:"followers"`
	Href            string          `json:"href"`
	ID              string          `json:"id"`
	Images          []Images        `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	URI             string          `json:"uri"`
}
type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}
type ExternalUrls struct {
	Spotify string `json:"spotify"`
}
type Followers struct {
	Href  interface{} `json:"href"`
	Total int         `json:"total"`
}
type Images struct {
	Height interface{} `json:"height"`
	URL    string      `json:"url"`
	Width  interface{} `json:"width"`
}
