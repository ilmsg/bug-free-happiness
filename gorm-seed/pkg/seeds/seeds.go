package seeds

import (
	"github.com/ilmsg/gorm-seed/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "Create User Jane",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Jane", 30)
			},
		},
		{
			Name: "Create User John",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "John", 30)
			},
		},
	}
}
