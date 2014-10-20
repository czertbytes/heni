package questions

import (
	"appengine"
	"appengine/datastore"
	"time"

	"github.com/czertbytes/heni/pkg/httputil"
	"github.com/czertbytes/heni/pkg/types"

	"github.com/mjibson/goon"
)

type Questions struct {
	n *goon.Goon
}

func NewQuestions(aeContext appengine.Context) *Questions {
	return &Questions{
		n: goon.FromContext(aeContext),
	}
}

func (q *Questions) Put(question *types.Question) (types.QuestionId, error) {
	now := time.Now()

	question.Status = types.QuestionStatusActive
	question.CreatedAt = now
	question.ModifiedAt = now
	question.SetFormattedValues()

	key, err := q.n.Put(question)
	if err != nil {
		return 0, httputil.InternalServerError("Oops! Something did not work as expected. Try it again, please.", err.Error(), err)
	}

	return types.QuestionId(key.IntID()), nil
}

func (q *Questions) Find(questionId types.QuestionId) (types.Question, error) {
	question := &types.Question{
		Id: questionId,
	}

	if err := q.n.Get(question); err != nil {
		var httpErr error
		if err == datastore.ErrNoSuchEntity {
			httpErr = httputil.NotFound("Requested question not found!", "", err)
		} else {
			httpErr = httputil.InternalServerError("Oops! Something did not work as expected. Try it again, please.", err.Error(), err)
		}

		return *question, httpErr
	}

	return *question, nil
}
