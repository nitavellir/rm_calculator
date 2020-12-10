package main

import (
	"flag"
	"fmt"
	"log"

	"rm_calculator/lib"
)

const (
	HIGH_REP = "high-rep"
	LOW_REP  = "low-rep"
)

var (
	kg       float64
	rep      int
	hard     bool
	targetRm int
	menu     string
)

func main() {
	flag.Float64Var(&kg, "kg", 0, "weight")
	flag.IntVar(&rep, "rep", 0, "reps")
	flag.BoolVar(&hard, "hard", false, "use hard mode")
	flag.IntVar(&targetRm, "rm", 0, "rm")
	flag.StringVar(&menu, "menu", "", "create menu(choose 'high-rep' or 'low-rep'")
	flag.Parse()

	if kg == 0 {
		log.Fatal("Specify the weight.")
	} else if rep == 0 {
		log.Fatal("Specify the reps.")
	}

	h := &lib.Handler{
		Kg:       kg,
		Rep:      rep,
		TargetRm: targetRm,
	}
	if hard {
		h.Correction = 30
	} else {
		h.Correction = 40
	}

	if err := h.CalcOneRm(); err != nil {
		log.Fatal("Can not calculate 1 rm.")
	}

	if hard {
		log.Println(fmt.Sprintf("1 RM(hard): %.1f", h.OneRm))
	} else {
		log.Println(fmt.Sprintf("1 RM: %.1f", h.OneRm))
	}

	if h.TargetRm > 0 {
		if kg, err := h.CalcRm(); err != nil {
			log.Fatal("Can not calculate the weight for the specified rm.")
		} else {
			log.Println(fmt.Sprintf("WEIGHT FOR %d RM: %.1f", h.TargetRm, kg))
		}
	}

	if menu == HIGH_REP {

	} else if menu == LOW_REP {

	}
}
