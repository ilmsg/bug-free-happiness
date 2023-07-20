package infar

import (
	"github.com/ilmsg/microauth/domain"
	"gorm.io/gorm"
)

type postgresStorage struct {
	*gorm.DB
}

type credential struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Password  string
	AuthToken domain.AuthToken `gorm:"embedded;embeddedPrefix:auth_token_"`
}

func toCredential(a *credential) domain.Credential {
	return domain.Credential{}
}

func fromCredential(a domain.Credential) *credential {
	return &credential{}
}

func NewPostgresStorage(db *gorm.DB) (domain.Storage, error) {
	if err := db.AutoMigrate(&credential{}); err != nil {
		return nil, err
	}
	return &postgresStorage{db}, nil
}
