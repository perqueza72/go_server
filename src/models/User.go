package models

type User struct {
	FullName string
	Email    string
	Age      uint8
}

type UserRequest struct {
	Name      string
	Lastname  string
	Full_name string
	Email     string
	Age       uint32
}
