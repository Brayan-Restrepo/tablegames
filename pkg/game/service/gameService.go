package service

import "tablegames/pkg/game/model"

type GameService interface {
	CreateGame(req *model.GameReq) (*model.Game, error)
	GetGameByID(id int) (*model.Game, error)
	DeleteGameByID(id int) error
}
