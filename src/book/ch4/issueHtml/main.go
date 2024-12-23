package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"book/packages"
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
	// заносим в переменную данные из файла по шаблону
	report := template.Must(template.New("issuehtml").Parse(issueList))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// обрабатываем корневой маршрут
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// выполняем шаблон
		if err := report.Execute(w, result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(r.Method)
		fmt.Println(r.URL.Query().Get("1"))
	})

	// обработчик маршрута /issues
	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		// получаем page из запроса
		pageStr := r.URL.Query().Get("page")
		// преобразуем page в цнло число
		page, err := strconv.Atoi(pageStr)
		// если параметр меньше 1, ставим значение по умолчанию
		if err != nil || page < 1 {
			page = 1 
		}
		
		// отправляем ответ с текущей страницей
		fmt.Fprintf(w, "Отображаем данные для страницы: %d", page)
	})
	
	
	http.ListenAndServe(":8080", nil)
}
