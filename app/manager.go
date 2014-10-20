package app

import (
	"github.com/czertbytes/heni/module/manager/question"
	"github.com/czertbytes/heni/module/manager/quiz"
	"github.com/czertbytes/heni/module/manager/quizcheck"

	"github.com/julienschmidt/httprouter"
)

func ManagerRoutes(router *httprouter.Router) {
	// Question
	router.POST("/manager/questions", question.Post)
	router.GET("/manager/questions/:questionId", question.Get)
	router.PUT("/manager/questions/:questionId", question.Put)
	router.PATCH("/manager/questions/:questionId", question.Patch)
	router.DELETE("/manager/questions/:questionId", question.Delete)

	// Quiz
	router.POST("/manager/quizes", quiz.Post)
	router.GET("/manager/quizes/:quizId", quiz.Get)
	router.PUT("/manager/quizes/:quizId", quiz.Put)
	router.PATCH("/manager/quizes/:quizId", quiz.Patch)
	router.DELETE("/manager/quizes/:quizId", quiz.Delete)

	// QuizCheck
	router.POST("/manager/quizes/:quizId/check", quizcheck.Post)

}
