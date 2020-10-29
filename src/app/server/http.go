package server

import (
	"net/http"
)


func ServeHTTP() error {

	Routing()
	var err = http.ListenAndServe(":8085", Router)
	if err != nil  {
		return err
	}

	return nil
}