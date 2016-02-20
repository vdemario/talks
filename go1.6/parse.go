package main

import (
	"fmt"
	"time"
)

func main() {
	const date = "2015-02-29T10:04:00+10:00"
	_, err := time.Parse(time.RFC3339, date)
	fmt.Println(err) // sem 29 de Fevereiro ano passado
}
