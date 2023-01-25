package custom_middlewares

import (
	"constants"
	"fmt"
	"log"
	"net/http"
	slice_methods "own_data_structures"
)

func UrlExistFunc() Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, exists := constants.URLS[r.URL.Path]; !exists {
				log.Default().Printf("Someone try to access to %v url", r.URL.Path)
				fmt.Fprintln(w, "404 Page not found.", http.StatusNotFound)

				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

func MethodAllowed() Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			methods, exists_path := constants.URLS[r.URL.Path]

			method := r.Method
			if exists := slice_methods.Find(methods, method); !exists || !exists_path {
				fmt.Fprintln(w, "This method is not allowed", http.StatusMethodNotAllowed)

				return
			}

			h.ServeHTTP(w, r)

		})
	}
}

// name      string
// 	lastname  string
// 	full_name string
// 	email     string
// 	age       uint32
