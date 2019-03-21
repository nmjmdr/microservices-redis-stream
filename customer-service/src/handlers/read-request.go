package handlers

import (
	"io/ioutil"
	"net/http"
)

// ReadRequest - Read a request to get the serialied data out
func ReadRequest(r *http.Request) ([]byte, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, err
}
