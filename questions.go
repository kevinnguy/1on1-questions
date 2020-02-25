package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type Question struct {
	Question string `json:"question"`
	Category string `json:"category"`
}

func main() {
	jsonFile, _ := os.Open("questions.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []Question
	json.Unmarshal(byteValue, &result)

	categories := make(map[string][]Question)
	for _, question := range result {
		category := question.Category
		categories[category] = append(categories[category], question)
	}

	for category, questions := range categories {
		randomInt := rand.Intn(len(questions))
		fmt.Println(randomInt, "Category:", category, "\nQuestion:", questions[randomInt].Question, "\n\n")
	}
}
