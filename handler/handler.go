package handler

import "github.com/jaloldinov/books-catalog/storage"

type handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{
		strg: strg,
	}
}
