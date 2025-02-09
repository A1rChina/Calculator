package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Calculation struct {
	EntryPrice  float64
	StopLoss    float64
	TakeProfit  float64
	RiskReward  float64
	StopLossPct float64
	TakeProfitPct float64
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	entryPrice, _ := strconv.ParseFloat(r.FormValue("entryPrice"), 64)
	stopLoss, _ := strconv.ParseFloat(r.FormValue("stopLoss"), 64)
	takeProfit, _ := strconv.ParseFloat(r.FormValue("takeProfit"), 64)

	riskReward := (takeProfit - entryPrice) / (entryPrice - stopLoss)
	stopLossPct := (entryPrice - stopLoss) / entryPrice * 100
	takeProfitPct := (takeProfit - entryPrice) / entryPrice * 100

	calculation := Calculation{
		EntryPrice:  entryPrice,
		StopLoss:    stopLoss,
		TakeProfit:  takeProfit,
		RiskReward:  riskReward,
		StopLossPct: stopLossPct,
		TakeProfitPct: takeProfitPct,
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, calculation)
}