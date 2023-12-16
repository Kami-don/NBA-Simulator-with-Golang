package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
}

type PlayerResponse struct {
	Data []struct {
		ID           int    `json:"id"`
		FirstName    string `json:"first_name"`
		HeightFeet   int    `json:"height_feet"`
		HeightInches int    `json:"height_inches"`
		LastName     string `json:"last_name"`
		Position     string `json:"position"`
		Team         Team   `json:"team"`
		WeightPounds int    `json:"weight_pounds"`
	} `json:"data"`
}

func main() {
	page := 1
	perPage := 100
	totalPages := 0

	for {
		url := fmt.Sprintf("https://www.balldontlie.io/api/v1/players?page=%d&per_page=%d", page, perPage)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("API isteği başarısız oldu:", err)
			break
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("API yanıtı başarısız:", resp.Status)
			break
		}

		// JSON verilerini oku ve parse et
		var playerResponse PlayerResponse
		err = json.NewDecoder(resp.Body).Decode(&playerResponse)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		if totalPages == 0 {
			totalPages = len(playerResponse.Data) / perPage
			if len(playerResponse.Data)%perPage != 0 {
				totalPages++
			}
		}

		players := playerResponse.Data
		for _, player := range players {
			fmt.Println(player.FirstName)
		}

		if page >= totalPages {
			break
		}

		page++
	}

}
