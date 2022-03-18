package model

type Score struct {
	Player string `json:"player" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Points int    `json:"points" binding:"required"`
}
