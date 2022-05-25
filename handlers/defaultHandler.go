package handlers

import "gorm.io/gorm"

type DefaultHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DefaultHandler {
	return DefaultHandler{db}
}
