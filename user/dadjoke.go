package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dadJokeResponse struct {
	id     string
	joke   string
	status int
}



func getRandomJoke() string {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var response dadJokeResponse
	json.Unmarshal(bodyBytes, &response)
	return response.joke
}
