package domain

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func NewUserResponseFromDomain(user User) UserResponse {
	return UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Balance: user.Balance,
	}
}
