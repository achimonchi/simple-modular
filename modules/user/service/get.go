package service

import "simple-modular/modules/user/domain"

func (s service) GetAll() (users []domain.User, err error) {
	return s.repo.FindAll()
}
