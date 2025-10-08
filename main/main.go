package main

import (
	// "bufio"
	"log"
	// "os"
	gui "gui.com/gui"
	// roulette "roulette.com/roulette"
)

func main() {

	log.SetPrefix("randomizer:")
	log.SetFlags(0)

	gui.NewProgram()

	// scanner := bufio.NewScanner(os.Stdin)

	// var value string

	// gui.Title()

	// for value != "spin" {

	// 	gui.Menu(roulette.GetOptions())

	// 	scanner.Scan()
	// 	value = scanner.Text()

	// 	if value == "q" {
	// 		os.Exit(1)
	// 	}

	// 	if value != "spin" {
	// 		roulette.AddElement(value)
	// 	}
	// }

	// result, err := roulette.Spin()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gui.Resultado(result)

}
