package types

import (
	"appengine/datastore"
	"time"
)

type QuestionId int64
type QuestionStatus int
type QuestionKind int

func ParseQuestionStatus(value int) QuestionStatus {
	switch value {
	case -1:
		return QuestionStatusDeleted
	case 1:
		return QuestionStatusActive
	default:
		return QuestionStatusUnknown
	}
}

func (qs QuestionStatus) String() string {
	switch qs {
	case QuestionStatusDeleted:
		return "deleted"
	case QuestionStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	QuestionStatusDeleted QuestionStatus = -1
	QuestionStatusUnknown QuestionStatus = 0
	QuestionStatusActive  QuestionStatus = 1
)

func ParseQuestionKind(value int) QuestionKind {
	switch value {
	case 1:
		return QuestionKindSingleChoice
	case 2:
		return QuestionKindMultipleChoice
	default:
		return QuestionKindUnknown
	}
}

func (qk QuestionKind) String() string {
	switch qk {
	case QuestionKindSingleChoice:
		return "single_choice"
	case QuestionKindMultipleChoice:
		return "multiple_choice"
	default:
		return "unknown"
	}
}

var (
	QuestionKindUnknown        QuestionKind = 0
	QuestionKindSingleChoice   QuestionKind = 1
	QuestionKindMultipleChoice QuestionKind = 2
)

type Question struct {
	Id              QuestionId     `json:"id,string" datastore:"-" goon:"id"`
	Status          QuestionStatus `json:"status"`
	StatusFormatted string         `json:"status_formatted"`
	CreatedAt       time.Time      `json:"created_at"`
	ModifiedAt      time.Time      `json:"modified_at"`

	Value         string       `json:"value"`
	Kind          QuestionKind `json:"kind"`
	KindFormatted string       `json:"kind_formatted"`
	Correct       []string     `json:"correct"`
	Option1       []string     `json:"option1,omitempty"`
	Option2       []string     `json:"option2,omitempty"`
	Option3       []string     `json:"option3,omitempty"`
	Option4       []string     `json:"option3,omitempty"`

	Parent *datastore.Key `json:"-" datastore:"-" goon:"parent"`
}

func (q *Question) SetFormattedValues() {
	q.SetStatusFormatted()
	q.SetKindFormatted()
}

func (q *Question) SetStatusFormatted() {
	q.StatusFormatted = q.Status.String()
}

func (q *Question) SetKindFormatted() {
	q.KindFormatted = q.Kind.String()
}
