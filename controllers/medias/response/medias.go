package response

import (
	"go-watchlist/business/medias"
	"time"
)

type MediaInsertResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Genres    string    `json:"genres"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainInsert(domain medias.Domain) MediaInsertResponse {
	return MediaInsertResponse{
		Message:   "Insert Media Succes",
		ID:        domain.ID,
		Type:      domain.Type,
		Title:     domain.Title,
		Genres:    domain.Genres,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
