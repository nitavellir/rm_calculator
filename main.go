package main

import (
	"flag"
	"fmt"
	"log"

	"rm_calculator/lib"
)

var (
	kg   float64
	rep  int
	hard bool
)

func main() {
	flag.Float64Var(&kg, "kg", 0, "weight")
	flag.IntVar(&rep, "rep", 0, "reps")
	flag.BoolVar(&hard, "hard", false, "use hard mode")
	flag.Parse()

	if kg == 0 {
		log.Fatal("Specify the weight.")
	} else if rep == 0 {
		log.Fatal("Specify the reps.")
	}

	h := &lib.Handler{
		Kg:       kg,
		Rep:      rep,
		HardMode: hard,
	}

	if err := h.CalcRm(); err != nil {
		log.Fatal("Can not calculate 1 rm.")
	}

	log.Println(fmt.Sprintf("1 RM: %.1f", h.OneRm))
}
