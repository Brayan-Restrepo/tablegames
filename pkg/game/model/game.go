package model

type (
	GameReq struct {
		Name          string `json:"name" validate:"required"`
		Avatar        string `json:"img" validate:"required"`
		NumberPlayers uint16 `json:"numberPlayers"`
		Ages          uint16 `json:"ages"`
	}
	Game struct {
		ID            uint64 `json:"id"`
		Name          string `json:"name" validate:"required"`
		Avatar        string `json:"img" validate:"required"`
		NumberPlayers uint16 `json:"numberPlayers"`
		Ages          uint16 `json:"ages"`
	}
)
