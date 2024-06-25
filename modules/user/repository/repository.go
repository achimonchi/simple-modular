package repository

import (
	"sync"
)

type repository struct {
	mut sync.Mutex
}

// karena kita butuh data yg konsisten, maka akan menggunakan pointer disini
func NewRepository() *repository {
	return &repository{}
}
