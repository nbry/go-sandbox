package castles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Booty struct {
	UserId       int    `json:"userId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	MissingBooty string `json:"missing"`
}

func makeRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	b := Booty{}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &b)

	fmt.Println(b.UserId)
	fmt.Println(b.MissingBooty)
}
