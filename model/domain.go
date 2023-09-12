package model

type Domain struct {
	Id     int `json:"id" binding:"required"`
	Domain string `json:"name" binding:"required"`
}