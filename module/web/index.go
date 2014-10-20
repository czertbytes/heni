package web

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const indexHTML = `
<!doctype html>
<html>
	<head>
		<title>Heni</title>
	</head>
	<body>
		<h1>Index</h1>
	</body>
</html>
`

var indexTemplate = template.Must(template.New("index").Parse(indexHTML))

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
