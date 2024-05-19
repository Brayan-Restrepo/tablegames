package entities

import (
	"tablegames/pkg/game/model"
	"time"
)

type Game struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;"`
	Name          string    `gorm:"type:varchar(128);not null;"`
	Avatar        string    `gorm:"type:varchar(256);not null;default:'';"`
	NumberPlayers uint16    `gorm:"type:bigint;not null;"`
	Ages          uint16    `gorm:"type:bigint;not null;"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime;"`
}

func (g *Game) ToGameModel() *model.Game {
	return &model.Game{
		ID:            g.ID,
		Name:          g.Name,
		NumberPlayers: g.NumberPlayers,
		Ages:          g.Ages,
	}
}
