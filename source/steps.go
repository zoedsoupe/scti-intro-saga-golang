package main

import "fmt"

func reserveHotel() error {
	fmt.Println("Reserva de hotel realizada. 💜")
	return nil
}

func cancelHotel() {
	fmt.Println("Reserva de hotel cancelada. 😭💔")
}

func runSteps() {
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

	fmt.Println("Todas as reservas realizadas com sucesso! 🎉")
}
