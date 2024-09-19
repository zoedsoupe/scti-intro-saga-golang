package main

import "fmt"

// Estrutura do SAGA
type Saga struct {
	steps         []func() error
	compensations []func()
}

// Adiciona um passo e sua compensação
func (s *Saga) AddStep(step func() error, compensation func()) {
	s.steps = append(s.steps, step)
	s.compensations = append(s.compensations, compensation)
}

// Executa o SAGA
func (s *Saga) Execute() error {
	for i, step := range s.steps {
		if err := step(); err != nil {
			// Executa as compensações em caso de erro
			for j := i - 1; j >= 0; j-- {
				s.compensations[j]()
			}
			return err
		}
	}
	return nil
}

func runManager() {
	saga := &Saga{}

	saga.AddStep(reserveFlight, cancelFlight)
	saga.AddStep(reserveHotel, cancelHotel)

	if err := saga.Execute(); err != nil {
		fmt.Println("Transação falhou:", err)
	} else {
		fmt.Println("Todas as reservas realizadas com sucesso! 🎉")
	}
}
