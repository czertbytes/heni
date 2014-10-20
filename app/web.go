package app

import (
	"github.com/czertbytes/heni/module/web"
	"github.com/czertbytes/heni/module/web/quiz"

	"github.com/julienschmidt/httprouter"
)

func WebRoutes(router *httprouter.Router) {
	router.GET("/", web.Get)

	router.GET("/web/quiz/:quizId", quiz.Get)
}
