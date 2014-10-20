package app

import (
	"log"
	"net/http"

	"github.com/czertbytes/heni/pkg/httputil/middleware"

	"github.com/julienschmidt/httprouter"
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	router := httprouter.New()

	WebRoutes(router)
	ManagerRoutes(router)

	http.Handle("/", middleware.NewLogging(router))
}
