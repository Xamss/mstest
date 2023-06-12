package delivery

import "net/http"

type delivery struct {
	Mux *http.ServeMux
}

func New() *delivery {
	mux := http.NewServeMux()
	return &delivery{Mux: mux}
}
