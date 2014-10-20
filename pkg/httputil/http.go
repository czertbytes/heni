package httputil

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	contentType := r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		dec := json.NewDecoder(r.Body)
		for {
			if err := dec.Decode(v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}

		return nil

	default:
		msg := fmt.Sprintf("Content type %s is not supported.", contentType)
		return BadRequest(msg, "", fmt.Errorf(msg))
	}

	return nil
}

func ResponseAsJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Fatal(err)
	}
}

func ResponseOK(w http.ResponseWriter, v interface{}) {
	ResponseAsJSON(w, http.StatusOK, v)
}

func ResponseError(w http.ResponseWriter, httpErr error) {
	ResponseAsJSON(w, httpErr.(*Error).StatusCode(), httpErr)
}

type Error struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	statusCode  int    `json:"-"`
	err         error  `json:"-"`
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func BadRequest(message, description string, err error) error {
	return &Error{
		Message:     message,
		Description: description,
		statusCode:  http.StatusBadRequest,
		err:         err,
	}
}

func NotFound(message, description string, err error) error {
	return &Error{
		Message:     message,
		Description: description,
		statusCode:  http.StatusNotFound,
		err:         err,
	}
}

func InternalServerError(message, description string, err error) error {
	return &Error{
		Message:     message,
		Description: description,
		statusCode:  http.StatusInternalServerError,
		err:         err,
	}
}
