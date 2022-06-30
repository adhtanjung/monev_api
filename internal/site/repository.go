package site

import (
	"fmt"

	"github.com/adhtanjung/monev/database"
	"github.com/adhtanjung/monev/internal/model"
)

type Repository interface {
	Create(site model.Site) error
}

type repository struct {
	db *database.Dbinstance
}

func NewSiteRepository(db *database.Dbinstance) Repository {
	return &repository{db: db}
}
func (r *repository) Create(site model.Site) error {

	// Check if username is already taken, if true return error
	if err := r.db.Db.Create(&site).Error; err != nil {
		return fmt.Errorf("failed to create site: %v", err)
	}
	return nil
}
