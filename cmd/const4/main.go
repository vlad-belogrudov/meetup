package main

import "fmt"

const smile_max_len = len("Биг Смайл")

func GetDefaultScore() map[string]int {
	return map[string]int{
		"Vasya": 33,
		"Petya": 44,
	}
}

func main() {
	score := GetDefaultScore()
	fmt.Println(score["Vasya"] + smile_max_len)
}
