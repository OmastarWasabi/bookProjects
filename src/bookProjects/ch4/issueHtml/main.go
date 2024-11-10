package main

import (
	"ch4/packages/github"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var issueList = `
<h1>{{.TotalCount}} тем </h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Titlr</th>
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
	report := template.Must(template.New("issuehtml").Parse(issueList))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := report.Execute(w, result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(r.Method)
		fmt.Println(r.URL.Query().Get("1"))
	})

	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1 // Если ошибка при преобразовании или номер страницы меньше 1
		}
	
		// Логика для получения данных для нужной страницы
		// Например, вы можете использовать "page" для выбора нужного набора данных
	
		// Выводим номер страницы
		fmt.Fprintf(w, "Отображаем данные для страницы: %d", page)
	})
	
	// http.ResponseWriter
	http.ListenAndServe(":8080", nil)
}
