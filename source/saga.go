package main

import "fmt"

type Saga struct {
	steps         []func() error // Passos da transação
	compensations []func()       // Operações compensatórias
}

func New() *Saga {
	return &Saga{}
}

func (s *Saga) AddStep(step func() error, compensation func()) {
	s.steps = append(s.steps, step)
	s.compensations = append(s.compensations, compensation)
}

func (s *Saga) Execute() error {
	for i, step := range s.steps {
		if err := step(); err != nil {
			for j := i - 1; j >= 0; j-- { // Desfazer etapas anteriores
				s.compensations[j]()
			}
			return err
		}
	}
	return nil
}

func runManager() {
	saga := New()

	saga.AddStep(reserveFlight, cancelFlight)
	saga.AddStep(reserveHotel, cancelHotel)
	saga.AddStep(reserveCar, cancelCar)

	if err := saga.Execute(); err != nil {
		fmt.Println("Transação falhou:", err)
	} else {
		fmt.Println("Transação realizada com sucesso.")
	}
}
