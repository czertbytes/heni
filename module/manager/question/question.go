package question

import (
	"appengine"

	"github.com/czertbytes/heni/pkg/questions"
	"github.com/czertbytes/heni/pkg/types"
)

func create(aeContext appengine.Context, question *types.Question) error {
	if _, err := questions.NewQuestions(aeContext).Put(question); err != nil {
		return err
	}

	return nil
}

func find(aeContext appengine.Context, id types.QuestionId) (types.Question, error) {
	return questions.NewQuestions(aeContext).Find(types.QuestionId(id))
}
