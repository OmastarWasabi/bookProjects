package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ComicsData struct {
	Num         int    `json:"num"`
	Title       string `json:"safe_title"`
	ImgURL      string `json:"img"`
	Description string `json:"transcript"`
}

func ComicsRequest(number string) (*ComicsData, error) {
	url := fmt.Sprintf("https://xkcd.com/%s/info.0.json", number)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сбой запроса: %s", err)
	}
	defer resp.Body.Close()
	var result ComicsData

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func ComicsWriteFile(data ComicsData) map[int]ComicsData {
	dataMap := make(map[int]ComicsData)

	fileContent, err := os.ReadFile("comicsList.json")
	if err != nil {
		fmt.Println("Ошибка четния JSON", err)

	}

	if err := json.Unmarshal(fileContent, &dataMap); err != nil {
		fmt.Println("Ошибка анмаршалинга", err)

	}

	dataMap[data.Num] = data
	dec, err := json.MarshalIndent(dataMap, "", " ")
	if err != nil {
		fmt.Println("Ошибка маршалинга", err)
		return dataMap

	}

	file, err := os.OpenFile("comicsList.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return dataMap

	}

	defer file.Close()

	_, err = file.Write(dec)
	if err != nil {
		fmt.Println("Ошибка записи", err)

	}
	return dataMap
}

