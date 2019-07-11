package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func getStory(path string) (map[string]Arc, error) {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	f := make(map[string]Arc)

	err = json.Unmarshal([]byte(content), &f)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func main() {

	story, err := getStory("story.json")

	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("story.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(t.Name())
		err := t.ExecuteTemplate(os.Stdout, "story.gohtml", story["intro"])
		if err != nil {
			fmt.Println(err)
		}
		err = t.ExecuteTemplate(rw, "story.gohtml", story["intro"])
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":8083", nil)
}
