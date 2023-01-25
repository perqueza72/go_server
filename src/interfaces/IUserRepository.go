package interfaces

import . "models"

type IUserRepository interface {
	Save(User) (User, error)
	Update(User) (User, error)
	Remove(User) (bool, error)
	Get(User) (User, error)
}
