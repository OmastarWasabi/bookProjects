package main

import (
	// "ch4/packages/github"
	// "html/template"
	// "os"
	// "fmt"
	"net/http"
)

var issueList = `
<h1>{{.TotalCount}} тем </h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr> 
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td> 
</tr>
{{end}}
</table>  
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "issuehtml.html")

	})
	http.ListenAndServe(":8080", nil)
}

