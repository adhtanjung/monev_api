package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type Site struct {
	BaseModel
	NamaLokasi string    `json:"nama_lokasi"`
	TerminalId string    `json:"terminal_id"`
	ProvinsiId uuid.UUID `gorm:"type:uuid" json:"provinsi_id"`
	Provinsi   Provinsi  `gorm:"foreignKey:ProvinsiID"`
}
