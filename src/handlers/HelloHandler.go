package handlersv1

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	handleGetHello(&w, r)
}

func handleGetHello(w *http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(*w, "Hello")
}
