package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type People struct {
	Name string `json:"name"`
}
type PeopleResponse struct {
	Number int      `json:"number"`
	People []People `json:"people"`
}

//	func main() {
//		text := `{"people":[{"craft":"ISS","number":1234567890,"name":"Sergey Rizhi"}]}`
//		textBytes := []byte(text)
//		people := People{}
//		err := json.Unmarshal(textBytes, &people)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(people.Number)
//
//		url := "http://api.open-notify.org/astros.json"
//		spaceClient := http.Client{
//			Timeout: time.Second * 2,
//		}
//
//		req, err := http.NewRequest(http.MethodGet, url, nil)
//		if err != nil {
//			log.Fatal(err)
//		}
//		req.Header.Set("User-Agent", "spacecount-tutorial")
//
//		res, err := spaceClient.Do(req)
//		if err != nil {
//			log.Fatal(err)
//		}
//		if res.Body != nil {
//			defer res.Body.Close()
//		}
//
//		fmt.Println("Http status: %s\n", res.Status)
//
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			log.Fatal(err)
//		}
//		p := People{}
//		err = json.Unmarshal(body, &p)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(p)
//
// }

func getAstros(apiURL string) (PeopleResponse, error) {
	p := PeopleResponse{}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return p, err
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}
	if res.Body == nil {

		defer res.Body.Close()
		return p, fmt.Errorf("no body in response")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}
	if err := json.Unmarshal(body, &p); err != nil {
		return p, err
	}

	return p, nil
}

func main() {
	apiURL := "http://api.open-notify.org/astros.json"
	response, err := getAstros(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range response.People {
		fmt.Printf("Let's wave to: %s\n", p.Name)
	}

}
