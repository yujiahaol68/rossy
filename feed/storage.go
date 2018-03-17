package feed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Title    string
	Link     string
	Desc     string
	Content  string
	Category string
	PubDate  string
	Status   string
}

type Source struct {
	Category     string `json:"category"`
	URL          string `json:"url"`
	ETag         string `json:"etag"`
	LastModified string `json:"last_modified"`
	Alias        string `json:"alias"`
	Type         string `json:"type"`
}

func SaveAsJSON(newSource []*Source, filePath string) error {
	if !hasSource(filePath) {
		return ioutil.WriteFile(filePath, toJSON(&newSource), 0644)
	}

	s, err := ReadExistSource(filePath)
	if err != nil {
		return err
	}

	repeatIndex := -1
	for _, oldSource := range s {
		for i, source := range newSource {
			if source.URL == oldSource.URL {
				repeatIndex = i
				break
			}
		}

		if repeatIndex != -1 {
			fmt.Printf("Found Repeat source: %s\n", newSource[repeatIndex].Alias)
			newSource = append(newSource[:repeatIndex], newSource[repeatIndex+1:]...)
			repeatIndex = -1
		}
	}

	if len(newSource) == 0 {
		return nil
	}

	s = append(s, newSource...)
	return ioutil.WriteFile(filePath, toJSON(&s), 0644)
}

func ReadExistSource(filePath string) ([]*Source, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var s []*Source
	err = json.Unmarshal(raw, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func hasSource(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func toJSON(p interface{}) []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return bytes
}
