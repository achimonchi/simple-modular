package repository

import (
	"sync"
)

type repository struct {
	mut sync.Mutex
}

func NewRepository() *repository {
	return &repository{}
}
