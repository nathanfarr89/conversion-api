package main

import (
	"conversion-api/tempconv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Temp struct {
	Number       float64             `json:"number"`
	FtCelsius    tempconv.Celsius    `json:"ftcelsius"`
	CtFahrenheit tempconv.Fahrenheit `json:"ctfahrenheit"`
	FtKelvin     tempconv.Kelvin     `json:"ftkelvin"`
	CtKelvin     tempconv.Kelvin     `json:"ctkelvin"`
	KtCelsius    tempconv.Celsius    `json:"ktcelsius"`
	KtFahrenheit tempconv.Fahrenheit `json:"ktfahrenheit"`
}

func returnTempConv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["id"]

	t, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "temp: %v\n", err)
		os.Exit(1)
	}

	ctf := tempconv.CToF(tempconv.Celsius(t))
	ftc := tempconv.FToC(tempconv.Fahrenheit(t))
	ftk := tempconv.FtK(tempconv.Fahrenheit(t))
	ctk := tempconv.CtK(tempconv.Celsius(t))
	ktf := tempconv.KtF(tempconv.Kelvin(t))
	ktc := tempconv.KtC(tempconv.Kelvin(t))
	conversion := Temp{
		Number:       t,
		FtCelsius:    ftc,
		CtFahrenheit: ctf,
		FtKelvin:     ftk,
		CtKelvin:     ctk,
		KtCelsius:    ktc,
		KtFahrenheit: ktf,
	}
	con, _ := json.Marshal(conversion)
	s := string(con)

	fmt.Fprintf(w, s)
}
