package main

import "fmt"

// Função para reservar voo
func reserveFlight() error {
	fmt.Println("Reserva de voo realizada. 💜")
	return nil
}

// Função para cancelar voo
func cancelFlight() {
	fmt.Println("Reserva de voo cancelada. 😭💔")
}

func runIntro() {
	err := reserveFlight()

	if err != nil {
		cancelFlight() // Operação compensatória
	}
}
