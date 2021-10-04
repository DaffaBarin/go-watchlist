package medias

import (
	"go-watchlist/business/medias"
)

type MediasCreate struct {
	Media_Id   string `json:"media_id"`
	Media_Type string `json:"media_type"`
}

type Medias struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	Title             string `json:"title"`
	Poster            string `json:"poster"`
	Overview          string `json:"overview"`
	Original_Language string `json:"original_language"`
	Genres            string `json:"genres"`
	Year              string `json:"year"`
	Duration          int    `json:"duration"`
}

func (req *Medias) ToDomain() *medias.Domain {
	return &medias.Domain{
		ID:                req.ID,
		Type:              req.Type,
		Title:             req.Title,
		Genres:            req.Genres,
		Poster:            req.Poster,
		Overview:          req.Overview,
		Original_Language: req.Original_Language,
		Year:              req.Year,
		Duration:          req.Duration,
	}
}
