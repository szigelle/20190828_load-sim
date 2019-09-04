//t, _ := time.Parse("2006-1-2 15:04", "2018-12-30 17:00")
package main

import (
	"20190828_load-sim/db"
	"20190828_load-sim/model"
	"20190828_load-sim/prep"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	year := 2018

	filename := "tacomaWA2013_2018"
	data := prep.ReadCSV(filename)
	train, test := prep.TrainTestSplit(data, year)

	n, h := prep.PullHoliday(train)
	ns, hs := prep.MakeSamples(n, h)

	////////////////////////////////////////////////////////////////////
	file, _ := os.Create(("output/test.csv"))
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	for i := range ns {
		x := []string{fmt.Sprintf("%d", i)}
		for j := range ns[i] {

			x = append(x, fmt.Sprintf("%s", ns[i][j].DATETIME.Format("2006-01-02 15:04")))

		}
		w.Write(x)
	}

	file, _ = os.Create(("output/htest.csv"))
	defer file.Close()
	w = csv.NewWriter(file)
	defer w.Flush()

	for i := range hs {
		x := []string{fmt.Sprintf("%s", i)}
		for j := range hs[i] {

			x = append(x, fmt.Sprintf("%s", hs[i][j].DATETIME.Format("2006-01-02 15:04")))

		}
		w.Write(x)
	}
	////////////////////////////////////////////////////////////////////

	nr, hr := model.BuildLR(ns, hs)

	///////////////////////////////////////
	n, h = prep.PullHoliday(test)

	nv, np, hv, hp := model.PredictLR(n, nr, h, hr)
	end := time.Now()

	runtime := end.Sub(start)
	fmt.Println("\nRUNTIME: ", runtime)

	fmt.Println("before (｡♥‿♥｡) ")
	fmt.Println("after ／(・x・)＼ ")

	writeResults(n, nv, np, h, hv, hp)
}

func writeResults(t map[time.Time]*db.DATA, tv map[int][]float64, tp map[int]float64, h map[time.Time]*db.DATA, hv map[string][]float64, hp map[string]float64) {
	file, _ := os.Create(fmt.Sprintf("output/%sresults%s.csv", "tacomaWA2013_2018", strconv.Itoa(2017)))
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write([]string{"ID", "HIGH", "LOW", "LOAD", "PREDICTION", "YDELTA", "msum", "HOUR", "DATETIME"})

	for i := range t {
		ID := t[i].ID
		w.Write([]string{
			strconv.Itoa(ID),
			fmt.Sprintf("%f", t[i].HIGH),
			fmt.Sprintf("%f", t[i].LOW),
			fmt.Sprintf("%f", t[i].LOAD),
			fmt.Sprintf("%f", tp[ID]),
			fmt.Sprintf("%f", (t[i].LOAD - tp[ID])),
			fmt.Sprintf("%f", (math.Abs(t[i].LOAD-tp[ID]))/t[i].LOAD),
			fmt.Sprintf("%d", t[i].HOUR),
			fmt.Sprintf("%s", t[i].DATETIME.Format("2006-01-02 15:04"))})
	}

	for i := range h {
		ID := h[i].HOLIDAY
		w.Write([]string{
			ID,
			fmt.Sprintf("%f", h[i].HIGH),
			fmt.Sprintf("%f", h[i].LOW),
			fmt.Sprintf("%f", h[i].LOAD),
			fmt.Sprintf("%f", hp[ID]),
			fmt.Sprintf("%f", (h[i].LOAD - hp[ID])),
			fmt.Sprintf("%f", (math.Abs(h[i].LOAD-hp[ID]))/h[i].LOAD),
			fmt.Sprintf("%d", h[i].HOUR),
			fmt.Sprintf("%s", h[i].DATETIME.Format("2006-01-02 15:04"))})
	}
}
