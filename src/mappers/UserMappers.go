package mappers

import (
	"models"
)

func UserRequestToModelsUser(user_data *models.UserRequest) *models.User {

	user := models.User{
		FullName: user_data.Name + " " + user_data.Lastname,
		Age:      uint8(user_data.Age),
		Email:    user_data.Email,
	}

	return &user

}
