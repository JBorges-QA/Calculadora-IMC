package main

import (
	"flag"
	"fmt"
	"os"
)

type Dados struct {
	peso   float64
	altura float64
}

func (d Dados) calcular() (string, error) {
	imc := d.peso / (d.altura * d.altura)
	var classificação string

	switch {
	case imc < 18.5:
		classificação = "Classificação: Magreza / Grau de Obesidade: 0"
		return classificação, nil
	case imc >= 18.5 && imc <= 24.9:
		classificação = "Classificação: Normal / Grau de Obesidade: 0"
		return classificação, nil
	case imc >= 25 && imc <= 29.9:
		classificação = "Classificação: Sobrepeso / Grau de Obesidade: 1"
		return classificação, nil
	case imc >= 30 && imc <= 39.9:
		classificação = "Classificação: Obesidade / Grau de Obesidade: 2"
		return classificação, nil
	case imc >= 40:
		classificação = "Classificação: Obesidade Grave / Grau de Obesidade: 3"
		return classificação, nil
	default:
		return "", fmt.Errorf("Erro no cálculo!")
	}
}

func main() {
	vPeso := flag.Float64("peso", 0, "O peso para o cálculo de IMC")
	vAltura := flag.Float64("altura", 0, "A altura para o cálculo de IMC")

	flag.Parse()

	IMC := Dados{
		peso:   *vPeso,
		altura: *vAltura,
	}

	resultado, err := IMC.calcular()
	if err != nil {
		fmt.Printf("Erro no cálculo: %v", err)
		os.Exit(1)
	}
	fmt.Printf(resultado)
}
