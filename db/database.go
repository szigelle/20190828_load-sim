package db

import "time"

//DATA is the orig data struct
type DATA struct {
	DATETIME time.Time // UTC
	LOCAL    time.Time // local time??
	YEAR     int
	MONTH    int
	ISOYR    int
	ISOWK    int
	WEEKDAY  int
	HOUR     int
	HOURSIN  float64
	HOURCOS  float64
	LOAD     float64
	HIGH     float64
	LOW      float64
	TDELTA   float64
	ID       int    //[ISOweek][Weekday][Hour]
	HOLIDAY  string // true if holiday
	BEFORE   bool
	AFTER    bool
}
