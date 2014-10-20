package quizcheck

import (
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager QuizCheck")
}
