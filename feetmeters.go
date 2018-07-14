package main

import (
	"conversion-api/feetmetersconv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type FeetMeters struct {
	Number float64               `json:"number"`
	Feet   feetmetersconv.Feet   `json:"feet"`
	Meters feetmetersconv.Meters `json:"meters"`
}

func returnFeetMetersConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["id"]

	t, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "feetmeters: %v\n", err)
		os.Exit(1)
	}

	f := feetmetersconv.MtF(feetmetersconv.Meters(t))
	m := feetmetersconv.FtM(feetmetersconv.Feet(t))
	conversion := FeetMeters{
		Number: t,
		Feet:   f,
		Meters: m,
	}
	con, _ := json.Marshal(conversion)
	s := string(con)

	fmt.Fprintf(w, s)
}
