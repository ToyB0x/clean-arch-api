package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func BuildErrorResponse(w http.ResponseWriter, str string) error {
	// err is original err, _err is another one that happen when fprintf
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _err := fmt.Fprintf(w, `{"result":"","error":%q}`, str)
	if _err != nil {
		log.Println("error in rendering middleware Error response: ", _err)
	}
	return _err
}
