package prep

import (
	"20190828_load-sim/db"
	"fmt"
	"time"
)

// GetData gets more data
func GetData(m map[time.Time]*db.DATA) {
	for i := range m {
		if t, ok := m[i.Add(-time.Hour)]; ok {
			m[i].PREVLOAD = t.LOAD
		}
		if t, ok := m[i.Add(time.Hour)]; ok {
			m[i].NEXTLOAD = t.LOAD
		}
	}
}

//TrainTestSplit splits the data
func TrainTestSplit(m map[time.Time]*db.DATA, year int) (map[time.Time]*db.DATA, map[time.Time]*db.DATA) {
	fmt.Println("...splitting data")
	train := make(map[time.Time]*db.DATA)
	test := make(map[time.Time]*db.DATA)

	for i := range m {

		if m[i].ISOYR == year {
			test[i] = m[i]
		} else {
			train[i] = m[i]
		}
	}
	return train, test
}

// MakeSamples map from intID -> array of DATA
func MakeSamples(t map[time.Time]*db.DATA, h map[time.Time]*db.DATA) (map[int][]*db.DATA, map[string][]*db.DATA) {

	ts := tsamples(t)
	hs := hsamples(h)

	return ts, hs
}

func tsamples(m map[time.Time]*db.DATA) map[int][]*db.DATA {
	//make ID idx
	s := make(map[int][]*db.DATA)

	for i := 1; i <= 52; i++ {
		for j := 0; j <= 6; j++ {
			for k := 0; k <= 23; k++ {
				var ID = i*10000 + j*100 + k
				x := make([]*db.DATA, 0)
				s[ID] = x
			}
		}
	}
	fmt.Println("...making samples for training")

	for j := range s {
		for i := range m {
			if m[i].ID == j { // ID matches the map id
				s[j] = append(s[j], m[i])
				if date, ok := m[i.AddDate(0, 0, -7)]; ok {
					if date.BEFORE == false { // if date is week before holiday
						s[j] = append(s[j], date)
					}
				}
				if date, ok := m[i.AddDate(0, 0, 7)]; ok {
					if date.AFTER == false { // if date is week after holiday
						s[j] = append(s[j], date)
					}
				}
			}
		}
	}
	return s
}

func hsamples(h map[time.Time]*db.DATA) map[string][]*db.DATA {
	//make ID idx
	s := make(map[string][]*db.DATA)

	hID := []string{"ny", "mlk", "dst", "mem", "indep", "labor", "thanks", "native", "eve", "xmas"}
	for i := range hID {
		for j := 0; j <= 23; j++ {
			x := make([]*db.DATA, 0)
			s[fmt.Sprintf("%s-%d", hID[i], j)] = x
		}
	}

	for j := range s {
		for i := range h {
			if h[i].HOLIDAY == j {

				s[j] = append(s[j], h[i])

			}
		}
	}

	return s
}
