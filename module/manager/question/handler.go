package question

import (
	"appengine"
	"log"
	"net/http"
	"strconv"

	"github.com/czertbytes/heni/pkg/httputil"
	"github.com/czertbytes/heni/pkg/types"

	"github.com/julienschmidt/httprouter"
)

func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aeContext := appengine.NewContext(r)

	var question types.Question
	if err := httputil.ParseBody(r, &question); err != nil {
		httputil.ResponseError(w, err)
		return
	}

	if err := create(aeContext, &question); err != nil {
		httputil.ResponseError(w, err)
		return
	}

	httputil.ResponseOK(w, question)
}

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aeContext := appengine.NewContext(r)

	questionIdStr := ps.ByName("id")
	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		err2 := httputil.BadRequest("Parsing request failed", err.Error(), err)
		httputil.ResponseError(w, err2)
		return
	}

	question, err := find(aeContext, types.QuestionId(questionId))
	if err != nil {
		httputil.ResponseError(w, err)
		return
	}

	httputil.ResponseOK(w, &question)
}

func Put(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager question")
}

func Patch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager question")
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager question")
}

func GetQuiz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Manager question")
}
