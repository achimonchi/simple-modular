package domain

type UserRequest struct {
	Name string `json:"name"`
}

func (u UserRequest) ToDomain() User {
	return User{
		Name:    u.Name,
		Balance: 0,
	}
}
