package domain

type User struct {
	Id      int
	Name    string
	Balance int
}

func (u User) IsExists() bool {
	return u.Id != 0
}
