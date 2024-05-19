package service

import (
	"tablegames/entities"
	"tablegames/pkg/game/model"
	_gameRepository "tablegames/pkg/game/repository"
)

type gameServiceImpl struct {
	gameRepository _gameRepository.GameRepository
}

func NewGameServiceImpl(gameRepository _gameRepository.GameRepository) GameService {
	return &gameServiceImpl{gameRepository}
}

func (s *gameServiceImpl) DeleteGameByID(id int) error {
	game, err := s.gameRepository.GetGameById(uint64(id), nil)
	if err != nil {
		return err
	}
	return s.gameRepository.DeleteGameById(game, nil)
}

func (s *gameServiceImpl) GetGameByID(id int) (*model.Game, error) {
	game, err := s.gameRepository.GetGameById(uint64(id), nil)
	if err != nil {
		return nil, err
	}

	return game.ToGameModel(), nil
}

func (s *gameServiceImpl) CreateGame(gameReq *model.GameReq) (*model.Game, error) {
	gameEntity := &entities.Game{
		Name:          gameReq.Name,
		Avatar:        gameReq.Avatar,
		NumberPlayers: gameReq.NumberPlayers,
		Ages:          gameReq.Ages,
	}
	game, err := s.gameRepository.CreateGame(gameEntity, nil)
	if err != nil {
		return nil, err
	}
	return game.ToGameModel(), nil
}
