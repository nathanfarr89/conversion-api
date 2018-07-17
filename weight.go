package main

import (
	"conversion-api/weightconv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Weight struct {
	Number    float64              `json:"number"`
	Pounds    weightconv.Pounds    `json:"pounds"`
	Kilograms weightconv.Kilograms `json:"kilograms"`
	Error     bool                 `json:"error"`
}

func returnWeightConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weight := vars["id"]

	t, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		conversion := Weight{
			Number:    0,
			Pounds:    0,
			Kilograms: 0,
			Error:     true,
		}
		c, _ := json.Marshal(conversion)
		s := string(c)

		io.WriteString(w, s)
	} else {
		k := weightconv.LbsToKgs(weightconv.Pounds(t))
		p := weightconv.KgsToLbs(weightconv.Kilograms(t))
		conversion := Weight{
			Number:    t,
			Pounds:    p,
			Kilograms: k,
			Error:     false,
		}
		c, _ := json.Marshal(conversion)
		s := string(c)

		io.WriteString(w, s)
	}

}
