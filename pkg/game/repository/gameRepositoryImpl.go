package repository

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"tablegames/databases"
	"tablegames/entities"
)

type gameRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewGameRepositoryImpl(db databases.Database, logger echo.Logger) GameRepository {
	return &gameRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *gameRepositoryImpl) GetGameById(id uint64, tx *gorm.DB) (*entities.Game, error) {
	var game entities.Game
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	if err := conn.First(&game, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("game not found with ID %d", id)
		}
		return nil, err
	}
	return &game, nil
}

func (r *gameRepositoryImpl) GetGames(tx *gorm.DB) (*[]entities.Game, error) {
	//TODO implement me
	panic("implement me")
	return nil, nil
}

func (r *gameRepositoryImpl) DeleteGameById(gameEntity *entities.Game, tx *gorm.DB) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	if err := conn.Delete(gameEntity).Error; err != nil {
		return err
	}
	return nil
}

func (r *gameRepositoryImpl) UpdateGame(gameEntity *entities.Game, tx *gorm.DB) (*entities.Game, error) {
	//TODO implement me
	panic("implement me")
	return nil, nil
}

func (r *gameRepositoryImpl) CreateGame(gameEntity *entities.Game, tx *gorm.DB) (*entities.Game, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	game := new(entities.Game)
	if err := conn.Create(gameEntity).Scan(game).Error; err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, err
	}
	return game, nil
}
