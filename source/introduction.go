package main

import "fmt"

// Definição das operações compensatórias
func reserveFlight() error {
	fmt.Println("Reserva de voo realizada.")
	return nil
}

func cancelFlight() {
	fmt.Println("Reserva de voo cancelada.")
}

// Função principal que inicia a transação SAGA
func runIntro() {
	err := reserveFlight()
	if err != nil {
		cancelFlight() // Operação compensatória
	}
}
