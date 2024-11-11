package main

import (
	"bookProjects/ch4/packages/github"
	"encoding/json"
	"os"
	

	"fmt"
)


func main (){
	if len(os.Args) < 5 {
		fmt.Println("Недостаточно аргуменотов. Введите: <owner><repos><token><method>")
		fmt.Println("                                                        <create>")
		fmt.Println("                                                        <update>")
		fmt.Println("                                                         <close>")
		
	} 

	owner := os.Args[1]
	repos := os.Args[2]
	token := os.Args[3]
	method := os.Args[4]

	
	var title string
	var body string
	switch method{
		
	
	case "create":
		
		fmt.Println("Введите: <title issue> и <body issue>")
		
		fmt.Scan(&title)
		fmt.Scan(&body)
		
		request := github.IssueRequest{
			Title: title,
			Body: body,
		}
		issue, err := github.CreateIssue(owner, repos, token, request)
		if err != nil{
			// fmt.Errorf("неверные параметры %v", err)
			return
		}
		fmt.Println(json.MarshalIndent(issue, "", " "))
	}
}

