package quiz

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const quizHTML = `
<!doctype html>
<html>
	<head>
		<title>Heni</title>
	</head>
	<body>
		<h1>Quiz</h1>
	</body>
</html>
`

var quizTemplate = template.Must(template.New("quiz").Parse(quizHTML))

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := quizTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
