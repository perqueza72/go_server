package custom_middlewares

import "net/http"

type Adapter func(http.HandlerFunc) http.HandlerFunc

func Adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		h = adapter(h)
	}

	return h
}
