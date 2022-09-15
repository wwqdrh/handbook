package cookbook

import (
	"fmt"
	"testing"
)

func TestUnix2Time(t *testing.T) {
	fmt.Println(Unix2Time(1659977921597))
}

func TestTimeDiff(t *testing.T) {
	fmt.Println(TimeDiff("2019-06-29", "2019-06-30"))
	fmt.Println(TimeDiff("2020-01-15", "2019-12-31"))
	fmt.Println(TimeDiff("2020-01-15", "1999-12-31"))
	fmt.Println(CustomTimeDiff("2019-06-29", "2019-06-30"))
	fmt.Println(CustomTimeDiff("2020-01-15", "2019-12-31"))
	fmt.Println(CustomTimeDiff("2020-01-15", "1999-12-31"))
}
