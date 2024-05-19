package repository

import (
	"gorm.io/gorm"
	"tablegames/entities"
)

type GameRepository interface {
	CreateGame(gameEntity *entities.Game, tx *gorm.DB) (*entities.Game, error)
	GetGameById(id uint64, tx *gorm.DB) (*entities.Game, error)
	GetGames(tx *gorm.DB) (*[]entities.Game, error)
	DeleteGameById(gameEntity *entities.Game, tx *gorm.DB) error
	UpdateGame(gameEntity *entities.Game, tx *gorm.DB) (*entities.Game, error)
}
