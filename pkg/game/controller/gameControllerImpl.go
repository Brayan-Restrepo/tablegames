package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tablegames/pkg/game/model"
	_gameService "tablegames/pkg/game/service"
)

type gameControllerImpl struct {
	gameService _gameService.GameService
}

func NewGameControllerImpl(gameService _gameService.GameService) GameController {
	return &gameControllerImpl{gameService: gameService}
}

func (c *gameControllerImpl) DeleteGameById(pctx echo.Context) error {
	idStr := pctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid game ID"})
	}
	if err := c.gameService.DeleteGameByID(id); err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return pctx.JSON(http.StatusOK, "Ok")
}

func (c *gameControllerImpl) GetGameByID(pctx echo.Context) error {

	idStr := pctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid game ID"})
	}
	game, err := c.gameService.GetGameByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pctx.JSON(http.StatusNotFound, map[string]string{"error": "Game not found"})
		}
		return pctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// Devolver la respuesta JSON con los datos del juego
	return pctx.JSON(http.StatusOK, game)
}

func (c *gameControllerImpl) CreateGame(pctx echo.Context) error {
	gameReq := new(model.GameReq)
	if err := pctx.Bind(gameReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	playerCoin, err := c.gameService.CreateGame(gameReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}
