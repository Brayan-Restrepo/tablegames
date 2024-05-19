package controller

import "github.com/labstack/echo/v4"

type GameController interface {
	CreateGame(pctx echo.Context) error
	GetGameByID(pctx echo.Context) error
	DeleteGameById(pctx echo.Context) error
}
