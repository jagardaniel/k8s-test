package handler

import "backend/model"

type Handler struct {
	DB map[int]*model.Person
}
