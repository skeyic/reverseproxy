package chain

import (
	"net/http"
)

type Filter interface {
	Handle(w http.ResponseWriter, req *http.Request) error
}
