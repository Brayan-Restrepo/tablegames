package server

import (
	"tablegames/pkg/game/controller"
	"tablegames/pkg/game/repository"
	"tablegames/pkg/game/service"
)

func (s *echoServer) initGameRouter() {
	router := s.app.Group("/v1/game")

	gameRepository := repository.NewGameRepositoryImpl(s.db, s.app.Logger)
	gameService := service.NewGameServiceImpl(gameRepository)
	gameController := controller.NewGameControllerImpl(gameService)

	router.POST("", gameController.CreateGame)
	router.GET("", gameController.GetGameByID)
	router.GET("/:id", gameController.GetGameByID)
	router.DELETE("/:id", gameController.DeleteGameById)
	router.PUT("/:id", gameController.CreateGame)
}
