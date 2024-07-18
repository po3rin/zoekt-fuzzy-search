package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
)

type ZoektResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	FileMatches []FileMatche `json:"FileMatches"`
}

type FileMatche struct {
	FileName string  `json:"FileName"`
	Repo     string  `json:"Repo"`
	Language string  `json:"Language"`
	Matches  []Match `json:"Matches"`
	URL      string  `json:"URL"`
}

type Match struct {
	LineNum   int        `json:"LineNum"`
	Fragments []Fragment `json:"Fragments"`
}

type Fragment struct {
	Pre   string `json:"Pre"`
	Match string `json:"Match"`
	Post  string `json:"Post"`
}

type Candidate struct {
	Name     string
	URL      string
	Fragment string
}

func runFuzzy(candidates []Candidate) error {
	i, err := fuzzyfinder.Find(
		candidates,
		func(i int) string {
			return candidates[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return candidates[i].Fragment
		}),
	)
	if err != nil {
		return err
	}
	fmt.Println(candidates[i].URL)
	return nil
}

func zoektResponse2Candidate(res ZoektResponse) ([]Candidate, error) {
	candidates := make([]Candidate, 0)
	for _, f := range res.Result.FileMatches {
		var fragment string
		for _, m := range f.Matches {
			for _, frag := range m.Fragments {
				fragment += color.HiBlueString(fmt.Sprint(m.LineNum)) + " : " + frag.Pre + color.YellowString(frag.Match) + frag.Post + "\n"
			}
		}
		candidates = append(candidates, Candidate{
			Name:     f.Repo + " : " + f.FileName,
			URL:      f.URL,
			Fragment: fragment,
		})
	}
	return candidates, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	data := ZoektResponse{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatal(err)
	}

	candidates, nil := zoektResponse2Candidate(data)
	if err != nil {
		log.Fatal(err)
	}

	err = runFuzzy(candidates)
	if err != nil {
		log.Fatal(err)
	}
}
