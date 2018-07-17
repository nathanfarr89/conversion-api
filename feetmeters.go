package main

import (
	"conversion-api/feetmetersconv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FeetMeters struct {
	Number float64               `json:"number"`
	Feet   feetmetersconv.Feet   `json:"feet"`
	Meters feetmetersconv.Meters `json:"meters"`
	Error  bool                  `json:"error"`
}

func returnFeetMetersConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["id"]

	t, err := strconv.ParseFloat(number, 64)
	if err != nil {
		conversion := FeetMeters{
			Number: 0,
			Feet:   0,
			Meters: 0,
			Error:  true,
		}
		con, _ := json.Marshal(conversion)
		s := string(con)

		io.WriteString(w, s)
	} else {
		f := feetmetersconv.MtF(feetmetersconv.Meters(t))
		m := feetmetersconv.FtM(feetmetersconv.Feet(t))
		conversion := FeetMeters{
			Number: t,
			Feet:   f,
			Meters: m,
			Error:  false,
		}
		con, _ := json.Marshal(conversion)
		s := string(con)

		io.WriteString(w, s)
	}

}
