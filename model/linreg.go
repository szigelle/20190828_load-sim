package model

import (
	"20190828_load-sim/db"
	"fmt"
	"time"

	"github.com/sajari/regression"
)

/*
func tLR(ts map[int][]*db.DATA) map[int]*regression.Regression {
	fmt.Println("...training tLR")
	tr := make(map[int]*regression.Regression)

	for i := range ts {
		var x1, x2, y []float64

		for j := range ts[i] {
			x1 = append(x1, ts[i][j].HIGH)
			x2 = append(x2, ts[i][j].LOW)
			y = append(y, ts[i][j].LOAD)
		}

		r := new(regression.Regression)
		r.SetObserved("load vs high")
		r.SetVar(0, "HIGH")
		r.SetVar(1, "LOW")
		for j := range y {
			r.Train(
				regression.DataPoint(y[j], []float64{x1[j], x2[j]}),
			)
		}
		r.Run()
		tr[i] = r
	}

	return tr
}

func hLR(hs map[string][]*db.DATA) map[string]*regression.Regression {

	fmt.Println("...training hLR")
	hr := make(map[string]*regression.Regression)

	for i := range hs {
		var x1, x2, y []float64

		for j := range hs[i] {
			x1 = append(x1, hs[i][j].HIGH)
			x2 = append(x2, hs[i][j].LOW)
			y = append(y, hs[i][j].LOAD)
		}

		r := new(regression.Regression)
		r.SetObserved("load vs high")
		r.SetVar(0, "HIGH")
		r.SetVar(1, "LOW")
		for j := range y {
			r.Train(
				regression.DataPoint(y[j], []float64{x1[j], x2[j]}),
			)
		}
		r.Run()
		hr[i] = r
	}

	return hr
}

//BuildLR builds the linear regression model
func BuildLR(t map[int][]*db.DATA, h map[string][]*db.DATA) (map[int]*regression.Regression, map[string]*regression.Regression) {

	tr := tLR(t)
	hr := hLR(h)

	return tr, hr
}

func tPred(t map[time.Time]*db.DATA, tr map[int]*regression.Regression) (map[int][]float64, map[int]float64) {
	s := make(map[int][]float64, 0)
	yhat := make(map[int]float64, 0)

	var x1, x2, y0 float64

	for i := range t {
		ID := t[i].ID
		x1 = t[i].HIGH
		x2 = t[i].LOW
		y0 = t[i].LOAD

		prediction, err := tr[ID].Predict([]float64{x1, x2})
		if err != nil {
			println(err)

		}
		yhat[ID] = prediction
		s[ID] = []float64{x1, x2, y0}
	}

	return s, yhat
}

func hPred(h map[time.Time]*db.DATA, hr map[string]*regression.Regression) (map[string][]float64, map[string]float64) {
	s := make(map[string][]float64, 0)
	yhat := make(map[string]float64, 0)

	for i := range h { // for all the test points
		ID := h[i].HOLIDAY
		s[ID] = []float64{h[i].HIGH, h[i].LOW}

		prediction, err := hr[ID].Predict([]float64{s[ID][0], s[ID][1]})

		if err != nil {
			println(err)

		}

		yhat[ID] = prediction
	}

	return s, yhat
}

// PredictLR predicts from regression
func PredictLR(t map[time.Time]*db.DATA, tr map[int]*regression.Regression, h map[time.Time]*db.DATA, hr map[string]*regression.Regression) (map[int][]float64, map[int]float64, map[string][]float64, map[string]float64) {
	tv, tp := tPred(t, tr)
	hv, hp := hPred(h, hr)
	return tv, tp, hv, hp
}
*/

//////////////////////////////////////////

func tLR(ts map[int][]*db.DATA) map[int]*regression.Regression {
	fmt.Println("...training tLR")
	tr := make(map[int]*regression.Regression)

	for i := range ts {
		var x1, x2, y []float64

		for j := range ts[i] {
			x1 = append(x1, ts[i][j].HIGH)
			x2 = append(x2, ts[i][j].LOW)

			y = append(y, ts[i][j].LOAD)
		}

		r := new(regression.Regression)
		r.SetObserved("load vs temp")
		r.SetVar(0, "h")
		r.SetVar(1, "l")

		for j := range y {
			r.Train(
				regression.DataPoint(y[j], []float64{x1[j], x2[j]}),
			)
		}
		r.Run()
		tr[i] = r
	}

	return tr
}

func hLR(hs map[string][]*db.DATA) map[string]*regression.Regression {

	fmt.Println("...training hLR")
	hr := make(map[string]*regression.Regression)

	for i := range hs {
		var x1, x2, y []float64

		for j := range hs[i] {
			x1 = append(x1, hs[i][j].HIGH)
			x2 = append(x2, hs[i][j].LOW)

			y = append(y, hs[i][j].LOAD)
		}

		r := new(regression.Regression)
		r.SetObserved("load vs temp")
		r.SetVar(0, "h")
		r.SetVar(1, "l")
		r.SetVar(2, "delta")

		for j := range y {
			r.Train(
				regression.DataPoint(y[j], []float64{x1[j], x2[j]}),
			)
		}
		r.Run()
		hr[i] = r
	}

	return hr
}

//BuildLR builds the linear regression model
func BuildLR(t map[int][]*db.DATA, h map[string][]*db.DATA) (map[int]*regression.Regression, map[string]*regression.Regression) {

	tr := tLR(t)
	hr := hLR(h)

	return tr, hr
}

func tPred(t map[time.Time]*db.DATA, tr map[int]*regression.Regression) (map[int][]float64, map[int]float64) {
	s := make(map[int][]float64, 0)
	yhat := make(map[int]float64, 0)

	var x1, x2, y0 float64

	for i := range t {
		ID := t[i].ID

		x1 = t[i].HIGH
		x2 = t[i].LOW

		y0 = t[i].LOAD

		prediction, err := tr[ID].Predict([]float64{x1, x2})
		if err != nil {
			println(err)

		}
		yhat[ID] = prediction
		s[ID] = []float64{x1, x2, y0}
	}

	return s, yhat
}

func hPred(h map[time.Time]*db.DATA, hr map[string]*regression.Regression) (map[string][]float64, map[string]float64) {
	s := make(map[string][]float64, 0)
	yhat := make(map[string]float64, 0)

	for i := range h { // for all the test points
		ID := h[i].HOLIDAY

		s[ID] = []float64{h[i].HIGH, h[i].LOW}

		prediction, err := hr[ID].Predict(s[ID])
		//		prediction, err := hr[ID].Predict([]float64{s[ID][0], s[ID][1], s[ID][2]})

		if err != nil {
			println(err)

		}

		yhat[ID] = prediction
	}

	return s, yhat
}

// PredictLR predicts from regression
func PredictLR(t map[time.Time]*db.DATA, tr map[int]*regression.Regression, h map[time.Time]*db.DATA, hr map[string]*regression.Regression) (map[int][]float64, map[int]float64, map[string][]float64, map[string]float64) {
	tv, tp := tPred(t, tr)
	hv, hp := hPred(h, hr)
	return tv, tp, hv, hp
}
