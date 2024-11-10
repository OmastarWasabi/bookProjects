package main


import (
    "time"
    "fmt"
    "log"
	"os"
	"ch4/packages/github"
)

                                                              // ИЗНАЧАЛЬНЫЙ main //
func main(){
	year, month, _ := time.Now().Date()
	YearNow := formatYear(year, month)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {
		year, month, _ := item.CreatedAt.Date()
		YearItem := formatYear(year, month)
		resString :=""
		if YearNow - YearItem < 1 {
			resString = "Менее месяца назад"
		} else {
			if YearNow - YearItem < 12 {
			resString = "Менее года назад"
			} else {resString = "Более года назад"}
		}

		fmt.Printf("#%-5d %9.9s %.55s давность: %s; %d %d\n",
			item.Number, item.User.Login, item.Title, resString, year, int(month))
	}
}

func formatYear(year int, month time.Month) int {
	YearNow := year *12 + int(month)
	return YearNow
}



// func main() {
//     owner := "OmastarWasabi"      // Замените на ваше имя пользователя или организацию
//     repo := "OmastarProject"      // Замените на ваше имя репозитория
//     issueNumber := 4                    // Номер issue, которую хотите закрыть
//     token := "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"                // Ваш токен доступа
// 	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", owner, repo, issueNumber)
// 	state := github.IssueRequest{
// 		State: "closed",
// 	}


//     res, err := github.SendRequest("PATCH", url, state, token)
// 	if err != nil {
//         fmt.Println("Ошибка:", err)
//     } else {
//         fmt.Println("Issue успешно закрыта.")
//     }
// 	total, _ := json.MarshalIndent(res, "", " ")
// 	fmt.Println(string(total))

// }
