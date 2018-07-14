package main

import (
	"conversion-api/measureconv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Measure struct {
	Number      float64                 `json:"number"`
	Inches      measureconv.Inches      `json:"inches"`
	Centimeters measureconv.Centimeters `json:"centimeters"`
}

func returnMeasureConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	measure := vars["id"]

	t, err := strconv.ParseFloat(measure, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "measure: %v\n", err)
		os.Exit(1)
	}

	i := measureconv.CmToI(measureconv.Centimeters(t))
	c := measureconv.IToCm(measureconv.Inches(t))
	conversion := Measure{
		Number:      t,
		Inches:      i,
		Centimeters: c,
	}
	con, _ := json.Marshal(conversion)
	s := string(con)

	fmt.Fprintf(w, s)
}
