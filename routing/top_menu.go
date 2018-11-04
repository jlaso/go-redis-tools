package routing

import (
	"log"
	"net/http"
	"text/template"
)


func TopMenuHandler(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var currentVal string
	current, ok := r.URL.Query()["current"]

	if !ok || len(current[0]) < 1 {
		log.Println("Url Param 'current' is missing")
		currentVal = "index.html"
	}else{
		currentVal = current[0]
	}

	tplt := `
	<div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom shadow-sm">
		<h5 class="my-0 mr-md-auto font-weight-normal">JLaso redis tools</h5>
		<nav class="my-2 my-md-0 mr-md-3" id="top-menu">
			<a class="p-2 text-dark {{if eq .Current "index.html"}}current{{end}}" href="/index.html">Home</a>
			<a class="p-2 text-dark {{if eq .Current "add-server.html"}}current{{end}}" href="/add-server.html">Add server</a>
			<a class="p-2 text-dark {{if eq .Current "servers.html"}}current{{end}}" href="/servers.html">Servers</a>
		</nav>
		<!-- <a class="btn btn-outline-primary" href="#">Sign up</a> -->
	</div>
`

	type tempStruct struct {
		Current string
	}
	a := tempStruct{currentVal}
	tmpl, err := template.New("test").Parse(tplt)
	if err != nil { panic(err) }

	err = tmpl.Execute(w, a)
	if err != nil { panic(err) }
}
