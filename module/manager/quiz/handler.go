package quiz

import (
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func Put(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func Patch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func GetQuestions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager quiz")
}

func PostCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager QuizCheck")
}
