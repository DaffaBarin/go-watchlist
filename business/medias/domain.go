package medias

import (
	"time"
)

type Domain struct {
	ID                int
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

// type Genres []struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

type Service interface {
	Create(domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id int) (Domain, error)
}

type Repository interface {
	Create(domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id int) (Domain, error)
}
