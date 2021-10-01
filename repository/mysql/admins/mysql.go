package admins

import (
	"context"
	"go-watchlist/business/admins"

	"gorm.io/gorm"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlAdminRepository(conn *gorm.DB) admins.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (rep *MysqlAdminRepository) Login(ctx context.Context, username string, password string) (admins.Domain, error) {
	var admin Admins
	result := rep.Conn.First(&admin, "username = ?", username)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return toDomain(admin), nil
}

func (rep *MysqlAdminRepository) Register(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	admin := fromDomain(domain)

	result := rep.Conn.Create(&admin)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return toDomain(admin), nil
}