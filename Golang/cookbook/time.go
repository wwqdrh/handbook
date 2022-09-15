package cookbook

import (
	"strconv"
	"strings"
	"time"
)

// 1659977921597
func Unix2Time(timestamp int64) string {
	return time.UnixMilli(timestamp).In(time.FixedZone("UTC+8", 8*3600)).Format("2006-01-02 15:04:05")
}

// 比较相差多少天
// yyyy-mm-dd
func TimeDiff(date1, date2 string) int {
	f := "2006-01-02"
	d1, _ := time.Parse(f, date1)
	d2, _ := time.Parse(f, date2)

	res := int(d1.Sub(d2).Hours() / 24)
	if res < 0 {
		return -res
	}
	return res
}

func CustomTimeDiff(date1, date2 string) int {
	return NewCustomDatefromStr(date1).SubDay(NewCustomDatefromStr(date2))
}

type customDate struct {
	year  int
	month int
	day   int
}

// yyyy-mm-dd
func NewCustomDatefromStr(str string) *customDate {
	parts := strings.Split(str, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	return &customDate{
		year:  year,
		month: month,
		day:   day,
	}
}

// 获取相差多少天
func (d *customDate) SubDay(other *customDate) int {
	if other.year > d.year ||
		(other.year == d.year && other.month > d.month) ||
		(other.year == d.year && other.month == d.month && other.day > d.day) {
		return other.SubDay(d)
	}

	if d.year-other.year > 1 {
		return other.SubDay(&customDate{year: other.year, month: 12, day: 31}) + (&customDate{
			year:  other.year + 1,
			month: 1,
			day:   0,
		}).SubDay(&customDate{
			year:  d.year - 1,
			month: 12,
			day:   31,
		}) + (&customDate{
			year:  d.year,
			month: 1,
			day:   0,
		}).SubDay(d)
	} else if d.year-other.year == 1 {
		return other.SubDay(&customDate{year: other.year, month: 12, day: 31}) + (&customDate{
			year:  d.year,
			month: 1,
			day:   0,
		}).SubDay(d)
	} else if d.month == other.month {
		return d.day - other.day
	} else {
		res := other.GetMonthDay(other.month) - other.day
		for i := other.month + 1; i < d.month; i++ {
			res += other.GetMonthDay(i)
		}
		res += d.day
		return res
	}
}

func (d *customDate) IsOddYear() bool {
	return (d.year%4 == 0 && d.year%100 != 0) || d.year%400 == 0
}

func (d *customDate) GetMonthDay(month int) int {
	if month == 2 && d.IsOddYear() {
		return 29
	}

	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	return months[month-1]
}
