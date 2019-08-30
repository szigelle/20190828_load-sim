package prep

import (
	"20190828_load-sim/db"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// ReadCSV if lalala
func ReadCSV(filename string) map[time.Time]*db.DATA {
	fmt.Println("...opening load csv file")
	f, err := os.Open(fmt.Sprintf("data/%s.csv", filename))
	if err != nil {
		log.Fatalln("COULD NOT OPEN CSV FILE", err)
		os.Exit(1)
	}

	fmt.Println("...reading csv")
	r := csv.NewReader(f)
	m := make(map[time.Time]*db.DATA)

	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}

	//	var x int = 0

	for { // pull data from each line
		line, err := r.Read()
		//		fmt.Println(x)
		//		x = x + 1
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		date := fmt.Sprintf("%s", line[1])
		high, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			fmt.Println(line[2])
			fmt.Println(err)
		}

		low, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			fmt.Println(line[3])
			fmt.Println(err)
		}

		for i := range line[4:] { // make a struct for each hour
			load, err := strconv.ParseFloat(line[(i+4)], 64)
			if err != nil {
				fmt.Println(line[i+4])
				fmt.Println(err)
			}

			hour := i // hour id

			datetime, err := time.Parse("1/2/2006 15:04", fmt.Sprintf("%s %d:00", date, hour))
			if err != nil {
				break
			}

			isoyr, isowk := datetime.ISOWeek()
			weekday := (int(datetime.Weekday()+6) % 7)

			if high == 0 || low == 0 || load == 0 {
				continue
			}

			m[datetime] = &db.DATA{
				DATETIME: datetime.UTC(),
				LOCAL:    datetime,
				YEAR:     datetime.Year(),
				MONTH:    int(datetime.Month()),
				ISOYR:    isoyr,
				ISOWK:    isowk,
				WEEKDAY:  weekday,
				HOUR:     datetime.Hour(),
				HOURSIN:  math.Sin(2 * math.Pi * float64(datetime.Hour()) / 23),
				HOURCOS:  math.Cos(2 * math.Pi * float64(datetime.Hour()) / 23),
				//https://www.kaggle.com/avanwyk/encoding-cyclical-features-for-deep-learning
				LOAD:    load,
				HIGH:    high,
				LOW:     low,
				TDELTA:  high - low,
				ID:      isowk*10000 + weekday*100 + datetime.Hour(), //[ISOweek][Weekday][Hour]
				HOLIDAY: "",
				BEFORE:  false,
				AFTER:   false,
			}
		}
	}
	return m
}

func toDAY(t time.Time) int {
	switch int(t.Weekday()) {
	case 0:
		return 7
	default:
		return (int(t.Weekday()))
	}
}
