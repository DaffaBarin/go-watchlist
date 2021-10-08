package medias

import (
	"go-watchlist/business/medias"

	"gorm.io/gorm"
)

type MysqlMediaRepository struct {
	Conn *gorm.DB
}

func NewMysqlMediaRepository(conn *gorm.DB) medias.Repository {
	return &MysqlMediaRepository{
		Conn: conn,
	}
}

func (rep *MysqlMediaRepository) Create(domain *medias.Domain) (medias.Domain, error) {
	media := fromDomain(*domain)
	if media.ID == 0 {
		return medias.Domain{}, nil
	}
	result := rep.Conn.Create(&media)

	if result.Error != nil {
		return medias.Domain{}, result.Error
	}

	return toDomain(media), nil
}

func (rep *MysqlMediaRepository) GetAll() ([]medias.Domain, error) {
	var media []Medias
	result := rep.Conn.Find(&media)

	if result.Error != nil {
		return nil, result.Error
	}

	return toListDomain(media), nil
}

func (rep *MysqlMediaRepository) GetByID(id int) (medias.Domain, error) {
	var media Medias
	result := rep.Conn.First(&media, "ID = ?", id)

	if result.Error != nil {
		return medias.Domain{}, result.Error
	}

	return toDomain(media), nil
}
