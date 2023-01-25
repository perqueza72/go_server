package handlersv1

import (
	"encoding/json"
	"fmt"
	"log"
	mappers "mappers_server"
	"models"
	"net/http"
	. "server_interfaces"
	"strconv"
)

type formHandler struct {
	userRepository IUserRepository
}

func NewFormHandler(userRepository IUserRepository) *formHandler {
	return &formHandler{
		userRepository: userRepository,
	}
}

func (f *formHandler) JsonFormHandler(w http.ResponseWriter, r *http.Request) {

	var userRequest = models.UserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		fmt.Fprintf(w, "Error getting Json request: %v", err)

		return
	}

	f.formPostHandler(w, &userRequest)
}

func (f *formHandler) FormStaticPageHandler(w http.ResponseWriter, r *http.Request) {
	defer log.Println("Request sended to formHandler")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error getting up data from static page. %v", err)
		return
	}

	var userRequest = models.UserRequest{
		Name:     r.FormValue("name"),
		Lastname: r.FormValue("last_name"),
		Email:    r.FormValue("email"),
	}
	userRequest.Full_name = userRequest.Name + " " + userRequest.Lastname
	age, _ := strconv.Atoi(r.FormValue("Age"))
	userRequest.Age = uint32(age)

	f.formPostHandler(w, &userRequest)

}

func (f *formHandler) formPostHandler(w http.ResponseWriter, userRequest *models.UserRequest) {

	user := mappers.UserRequestToModelsUser(userRequest)
	if _, err := f.userRepository.Save(*user); err != nil {
		fmt.Fprintf(w, "Error obtaining data. %v", err)
		return
	}

	fmt.Fprintf(w, "Succesfully handled. %v\n", *user)
}
