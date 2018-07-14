package main

import (
	"conversion-api/weightconv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Weight struct {
	Number    float64              `json:"number"`
	Pounds    weightconv.Pounds    `json:"pounds"`
	Kilograms weightconv.Kilograms `json:"kilograms"`
}

func returnWeightConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weight := vars["id"]

	t, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "weight: %v\n", err)
		os.Exit(1)
	}

	k := weightconv.LbsToKgs(weightconv.Pounds(t))
	p := weightconv.KgsToLbs(weightconv.Kilograms(t))
	conversion := Weight{
		Number:    t,
		Pounds:    p,
		Kilograms: k,
	}
	c, _ := json.Marshal(conversion)
	s := string(c)

	fmt.Fprintf(w, s)
}
