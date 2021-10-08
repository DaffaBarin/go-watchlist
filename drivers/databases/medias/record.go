package medias

import (
	"go-watchlist/business/medias"
	"time"
)

type Medias struct {
	ID                int `gorm:"primaryKey"`
	Type              string
	Title             string
	Poster            string
	Overview          string
	Original_Language string
	Genres            string
	Year              string
	Duration          int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func toDomain(media Medias) medias.Domain {
	return medias.Domain{
		ID:                media.ID,
		Type:              media.Type,
		Title:             media.Title,
		Poster:            media.Poster,
		Overview:          media.Overview,
		Original_Language: media.Original_Language,
		Genres:            media.Genres,
		Year:              media.Year,
		Duration:          media.Duration,
		CreatedAt:         media.CreatedAt,
		UpdatedAt:         media.UpdatedAt,
	}
}

func toListDomain(domain []Medias) (result []medias.Domain) {
	result = []medias.Domain{}
	for _, domain := range domain {
		result = append(result, toDomain(domain))
	}
	return result
}

func fromDomain(domain medias.Domain) Medias {
	return Medias{
		ID:                domain.ID,
		Type:              domain.Type,
		Title:             domain.Title,
		Poster:            domain.Poster,
		Overview:          domain.Overview,
		Original_Language: domain.Original_Language,
		Genres:            domain.Genres,
		Year:              domain.Year,
		Duration:          domain.Duration,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
