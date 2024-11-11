package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"bytes"
    
)

const IssuesURL = "https://api.github.com/search/issues"

// Структура, которая хранит кол-во проблем и сами проблемы
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Структура, которая хранит параметры проблемы
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// Данные о пользователе из структуры проблем
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueRequest struct {
	Title string `json:"title"`
	Body string ` json:"body,omitempty"`
	State string `json:"state,omitempty"`
	
}

// Поиск проблем
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " ")) // Join делает из среза строку с разделителем " ",
	// Query преобразует символы в строке в формат URL
	resp, err := http.Get(IssuesURL + "?q=" + q) // Get выполняет запрос по URL и если получает ответ возвращает resp
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK { // смотрим в resp и сравниваем с константой
		return nil, fmt.Errorf("cбой запроса: %s", resp.Status)
	}
	var result IssuesSearchResult                                      // Создаем переменную
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // NewDecoder читает данные из resp.Body и декодирует в result
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}


func CreateIssue(owner, repo, token string, request IssueRequest) (*Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	return SendRequest("POST", url, request, token)
}

// Обновление существующей проблемы
func UpdateIssue(owner, repo string, issueNumber int, token string, request IssueRequest) (*Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", owner, repo, issueNumber)
	return SendRequest("PATCH", url, request, token)
}

// Закрытие проблемы
func CloseIssue(owner, repo string, issueNumber int, token string) (*Issue, error) {
	request := IssueRequest{State: "closed"}
	return UpdateIssue(owner, repo, issueNumber, token, request)
}

// Вспомогательная функция для отправки запросов POST и PATCH
func SendRequest(method, url string, request IssueRequest, token string) (*Issue, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("ошибка запроса: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}


