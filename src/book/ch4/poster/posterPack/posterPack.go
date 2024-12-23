package posterPack

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

type FilmData struct{
	Title string `json:"Title"`
	Poster string  `json:"Poster"`
}

type SearchData struct {
	Search []FilmData `json:"Search"`
}
func PosterRequest(film string) (*SearchData, error){
	url := fmt.Sprintf("https://omdbapi.com/?apikey=748dbac3&s=%s", film)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("ошибка запроса %s", err)
	}

	defer resp.Body.Close()
	var request SearchData

	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		return nil, err
	}

	return &request, err


}

