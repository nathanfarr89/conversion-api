package main

import (
	"conversion-api/measureconv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Measure struct {
	Number      float64                 `json:"number"`
	Inches      measureconv.Inches      `json:"inches"`
	Centimeters measureconv.Centimeters `json:"centimeters"`
	Error       bool                    `json:"error"`
}

func returnMeasureConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	measure := vars["id"]

	t, err := strconv.ParseFloat(measure, 64)
	if err != nil {
		conversion := Measure{
			Number:      0,
			Inches:      0,
			Centimeters: 0,
			Error:       true,
		}
		con, _ := json.Marshal(conversion)
		s := string(con)

		io.WriteString(w, s)
	} else {
		i := measureconv.CmToI(measureconv.Centimeters(t))
		c := measureconv.IToCm(measureconv.Inches(t))
		conversion := Measure{
			Number:      t,
			Inches:      i,
			Centimeters: c,
			Error:       false,
		}
		con, _ := json.Marshal(conversion)
		s := string(con)

		io.WriteString(w, s)
	}

}
