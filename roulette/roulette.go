package roulette

//make faz um slice
//len para tamanho, cap para capacidade
//operador : mudar de tamanho dinamicamente (limite no cap)
//copy para copiar slice para outro

import (
	"errors"
	"math/rand"
)

var options = make([]string, 0, 10)

func addElement(option string) error {

	if option == "" {
		return errors.New("option is empty")
	}

	if len(options) == cap(options) {
		increaseCap(&options)
	}

	options = append(options, option)

	return nil
}

func spin() (string, error) {

	if len(options) == 0 {
		return "", errors.New("options are empty")
	}

	element := options[rand.Intn(len(options))]

	return element, nil

}

func removeElement(value string) {

	for i, v := range options {
		if v == value {
			options = append(options[:i], options[i+1:]...)
			break
		}
	}

}

// go já gerencia o aumento do cap pelo append, função inutil
func increaseCap(options *[]string) {

	optionsNew := make([]string, len(*options), cap(*options)+10)
	copy(optionsNew, *options)

	*options = optionsNew
}
