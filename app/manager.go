package app

import (
	"github.com/czertbytes/heni/module/manager/question"
	"github.com/czertbytes/heni/module/manager/quiz"

	"github.com/julienschmidt/httprouter"
)

func ManagerRoutes(router *httprouter.Router) {
	// Question
	router.POST("/manager/questions", question.Post)
	router.GET("/manager/questions/:id", question.Get)
	router.PUT("/manager/questions/:id", question.Put)
	router.PATCH("/manager/questions/:id", question.Patch)
	router.DELETE("/manager/questions/:id", question.Delete)
	router.GET("/manager/questions/:id/quiz", question.GetQuiz)

	// Quiz
	router.POST("/manager/quizes", quiz.Post)
	router.GET("/manager/quizes/:id", quiz.Get)
	router.PUT("/manager/quizes/:id", quiz.Put)
	router.PATCH("/manager/quizes/:id", quiz.Patch)
	router.DELETE("/manager/quizes/:id", quiz.Delete)
	router.GET("/manager/quizes/:id/questions", quiz.GetQuestions)
	router.POST("/manager/quizes/:id/check", quiz.PostCheck)
}
