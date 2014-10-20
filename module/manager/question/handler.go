package question

import (
	"appengine"
	"log"
	"net/http"
	"strconv"

	"github.com/czertbytes/heni/pkg/httputil"
	"github.com/czertbytes/heni/pkg/questions"
	"github.com/czertbytes/heni/pkg/types"

	"github.com/julienschmidt/httprouter"
)

func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aeContext := appengine.NewContext(r)

	var question types.Question
	if httpErr := httputil.ParseBody(r, &question); httpErr != nil {
		httputil.ResponseError(w, httpErr)
		return
	}

	if _, httpErr := questions.NewQuestions(aeContext).Put(&question); httpErr != nil {
		httputil.ResponseError(w, httpErr)
		return
	}

	httputil.ResponseOK(w, question)
}

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aeContext := appengine.NewContext(r)

	questionIdStr := ps.ByName("questionId")
	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		httpErr := httputil.BadRequest("Parsing request failed", err.Error(), err)
		httputil.ResponseError(w, httpErr)
		return
	}

	question, httpErr := questions.NewQuestions(aeContext).Find(types.QuestionId(questionId))
	if httpErr != nil {
		httputil.ResponseError(w, httpErr)
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
