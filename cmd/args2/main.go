package main

import (
	"encoding/json"
	"fmt"
)

type Score struct {
	Team  string
	Score int
}

func main() {
	var score Score
	input := `{"Team": "Yadro", "score": 777}`

	json.Unmarshal([]byte(input), &score)
	fmt.Println(score)
}
