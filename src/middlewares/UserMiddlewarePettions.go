package custom_middlewares

import (
	"fmt"
	"models"
	"reflect"
)

func Main() {
	s := models.UserRequest{
		Name: "Chetan", Lastname: "Kumar", Full_name: "Bangalore", Email: "email", Age: 123}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\n", typeOfS.Field(i).Name)
	}
}
