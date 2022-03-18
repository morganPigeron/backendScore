package model

type Score struct {
	Player string `json:"user" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Points int    `json:"points" binding:"required"`
}
