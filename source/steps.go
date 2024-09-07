package main

import "fmt"

func reserveHotel() error {
	fmt.Println("Reserva de hotel realizada.")
	return nil
}

func cancelHotel() {
	fmt.Println("Reserva de hotel cancelada.")
}

func reserveCar() error {
	fmt.Println("Reserva de carro realizada.")
	return nil
}

func cancelCar() {
	fmt.Println("Reserva de carro cancelada.")
}

func processSaga() {
	err := reserveFlight()
	if err != nil {
		cancelFlight()
		return
	}

	err = reserveHotel()
	if err != nil {
		cancelFlight()
		cancelHotel()
		return
	}

	err = reserveCar()
	if err != nil {
		cancelFlight()
		cancelHotel()
		cancelCar()
	}
}

func runSteps() {
	processSaga() // Executa a transação SAGA completa
}
