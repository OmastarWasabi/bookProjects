package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"book/packages"
)

const templ = `{{.TotalCount}} тем:
	{{range .Items}}
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`

// daysAgo возвращает количество дней, прошедших с даты t
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	// инициализация шаблона с функцией daysAgo
	report := template.Must(template.New("issueReport").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	// выполняем поиск по GitHub Issues
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// выполняем шаблон с выводом результата
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
