package main

import (
	"constants"
	middleware "custom_middlewares"
	"fmt"
	handlers "handlersv1"
	"log"
	"net/http"
	user_repository "repository"
)

func main() {

	defer log.Println("Server shout down.")

	fmt.Printf("Starting server on port %v\n", constants.PORT)

	user_repository := user_repository.UserRepository{}
	formHandler := handlers.NewFormHandler(user_repository)

	http.HandleFunc(constants.PAGES[constants.HELLO_PAGE],
		middleware.Adapt(handlers.HandleHello,
			middleware.UrlExistFunc(),
			middleware.MethodAllowed()))

	http.HandleFunc(constants.PAGES[constants.FORM_PAGE],
		middleware.Adapt(formHandler.JsonFormHandler,
			middleware.UrlExistFunc(),
			middleware.MethodAllowed()))

	if err := http.ListenAndServe(":"+constants.PORT, nil); err != nil {
		log.Fatal(err)
	}

}
