package types

import (
	"time"
)

type QuizId int64
type QuizStatus int

func ParseQuizStatus(value int) QuizStatus {
	switch value {
	case -1:
		return QuizStatusDeleted
	case 1:
		return QuizStatusActive
	default:
		return QuizStatusUnknown
	}
}

func (qs QuizStatus) String() string {
	switch qs {
	case QuizStatusDeleted:
		return "deleted"
	case QuizStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	QuizStatusDeleted QuizStatus = -1
	QuizStatusUnknown QuizStatus = 0
	QuizStatusActive  QuizStatus = 1
)

type Quiz struct {
	Id              QuizId     `json:"id,string" datastore:"-" goon:"id"`
	Status          QuizStatus `json:"status"`
	StatusFormatted string     `json:"status_formatted"`
	CreatedAt       time.Time  `json:"created_at"`
	ModifiedAt      time.Time  `json:"modified_at"`

	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (q *Quiz) SetFormattedValues() {
	q.SetStatusFormatted()
}

func (q *Quiz) SetStatusFormatted() {
	q.StatusFormatted = q.Status.String()
}
