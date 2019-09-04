package prep

//New Year's Day, Independence Day, Veterans Day, and Christmas Day are observed
//on the same calendar date each year. Holidays that fall on a Saturday are observed
//by federal employees who work a standard Monday to Friday week on the previous Friday.
//Federal employees who work on Saturday will observe the holiday on Saturday; Friday will
//be a regular work day. Holidays that fall on a Sunday are observed by federal workers the following Monday.
//The other holidays always fall on a particular day of the week.[5]
// SHIFT .Day() s.t. Monday = 1, Sunday = 7 for calculation purposes

import (
	"20190828_load-sim/db"
	"fmt"
	"math"
	"time"
)

func compareDates(d *db.DATA, wom int, dow int) bool {

	f := time.Date(d.LOCAL.Year(), d.LOCAL.Month(), 1, 0, 0, 0, 0, time.UTC) // !!!change utc to pass local
	f = f.AddDate(0, 0, 7*(wom-1))

	dof := toDAY(f)
	delta := int(math.Abs(float64(dof) - float64(dow)))

	if dof < dow {
		f = f.AddDate(0, 0, delta)
	} else if dof > dow {
		f = f.AddDate(0, 0, (7 - delta))
	}

	if f.Day() == d.LOCAL.Day() {
		return true
	}
	return false
}

// Formulaic Holidays
func mlk(d *db.DATA) bool {
	weekofM := 3
	dayofW := 1 //R
	return compareDates(d, weekofM, dayofW)
}

func sf(d *db.DATA) bool {
	weekofM := 2
	dayofW := 7
	return compareDates(d, weekofM, dayofW)
}

func mem(d *db.DATA) bool {
	weekofM := 4
	// account for fifth week
	if toDAY(time.Date(d.LOCAL.Year(), d.LOCAL.Month(), 1, 0, 0, 0, 0, time.UTC)) > 5 {
		weekofM = 5
	}
	dayofW := 1 //M
	return compareDates(d, weekofM, dayofW)
}

func labor(d *db.DATA) bool {
	weekofM := 1
	dayofW := 1 //M
	return compareDates(d, weekofM, dayofW)
}
func fb(d *db.DATA) bool {
	weekofM := 1

	dayofW := 7
	return compareDates(d, weekofM, dayofW)
}

func thanks(d *db.DATA) bool {
	weekofM := 4
	dayofW := 4
	return compareDates(d, weekofM, dayofW)
}

func native(d *db.DATA) bool {
	weekofM := 4
	dayofW := 5
	return compareDates(d, weekofM, dayofW)
}

//PullHoliday is a holiday helper
func PullHoliday(m map[time.Time]*db.DATA) (map[time.Time]*db.DATA, map[time.Time]*db.DATA) {
	fmt.Println("...pulling holidays")
	h := make(map[time.Time]*db.DATA)

	for i := range m {
		switch m[i].LOCAL.Month() {
		case 1:
			switch m[i].LOCAL.Day() {
			case 1:
				m, h = isHoliday(m, h, i, "ny") // new year
				continue
			default:
				if mlk(m[i]) == true {
					m, h = isHoliday(m, h, i, "mlk")
					continue
				}
			}
		case 2:

		case 3:
			switch m[i].LOCAL.Day() {
			default:
				if sf(m[i]) == true {
					m, h = isHoliday(m, h, i, "dst")
					continue
				}
			}
		//case 4:
		case 5:
			switch m[i].LOCAL.Day() {
			default:
				if mem(m[i]) == true {
					m, h = isHoliday(m, h, i, "mem")
					continue
				}
			}
		//case 6:
		case 7:
			switch m[i].LOCAL.Day() {
			case 4:
				m, h = isHoliday(m, h, i, "indep") // independence day
				continue
			}
		//case 8:
		case 9:
			switch m[i].LOCAL.Day() {
			default:
				if labor(m[i]) == true {
					m, h = isHoliday(m, h, i, "labor")
					continue
				}
			}

		//case 10:
		case 11:
			switch m[i].LOCAL.Day() {
			default:
				switch thanks(m[i]) {
				case true:
					m, h = isHoliday(m, h, i, "thanks")
					continue
				default:
					if native(m[i]) == true { // day after thanksgiving
						m, h = isHoliday(m, h, i, "native")
						continue
					}
				}
			}
		case 12:
			switch m[i].LOCAL.Day() {
			case 24: // christmas eve
				m, h = isHoliday(m, h, i, "eve")
				continue

			case 25: // christmas
				m, h = isHoliday(m, h, i, "xmas")
				continue
			}

		}
	}
	return m, h
}

// if key.isHoliday, then delete key from map
func isHoliday(m map[time.Time]*db.DATA, h map[time.Time]*db.DATA, i time.Time, s string) (map[time.Time]*db.DATA, map[time.Time]*db.DATA) {
	m[i].HOLIDAY = fmt.Sprintf("%s-%d", s, m[i].HOUR)

	h[i] = m[i]

	start := i.AddDate(0, 0, -7)
	end := i

	//set before and after
	for i := start; i == end; i.AddDate(0, 0, 1) {
		m[i].BEFORE = true
	}

	start = i
	end = i.AddDate(0, 0, 7)

	for i := start; i == end; i.AddDate(0, 0, 1) {
		m[i].AFTER = true
	}

	delete(m, i)
	return m, h
}
