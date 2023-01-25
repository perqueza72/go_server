package user_repository

import . "models"

type UserRepository struct{}

func (UserRepository UserRepository) Save(user User) (User, error) {
	return user, nil
}

func (UserRepository UserRepository) Update(user User) (User, error) {
	return user, nil
}

func (UserRepository UserRepository) Remove(user User) (bool, error) {
	return false, nil
}

func (UserRepository UserRepository) Get(user User) (User, error) {
	return user, nil
}
