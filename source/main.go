package main

import "flag"

func main() {
	intro := flag.Bool("intro", false, "Run introduction")
	steps := flag.Bool("steps", false, "Run multi step example")
	manager := flag.Bool("manager", false, "Run manager example")

	flag.Parse()

	if *intro {
		runIntro()
	}

	if *steps {
		runSteps()
	}

	if *manager {
		runManager()
	}
}
