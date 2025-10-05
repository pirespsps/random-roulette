package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	roulette "roulette.com/roulette"
)

func main() {

	log.SetPrefix("randomizer:")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	var value string

	for value != "spin" {

		fmt.Print(roulette.PrintOptions())

		fmt.Print("Digite o valor | \"spin\" para rodar\n")
		scanner.Scan()

		value = scanner.Text()

		fmt.Print("\n-----------------------------------------------\n")

		if value != "spin" {
			roulette.AddElement(value)
		}
	}

	result, err := roulette.Spin()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Resultado: %v", result)

}
