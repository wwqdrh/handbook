package os

import (
	"fmt"
	"time"
)

func TimeExample() {
	p := fmt.Println

	// 从获取当前时间时间开始。
	now := time.Now()
	p(now)

	// 通过提供年月日等信息，你可以构建一个 `time`。
	// 时间总是与 `Location` 有关，也就是时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 支持通过 `Weekday` 输出星期一到星期日。
	p(then.Weekday())

	// 这些方法用来比较两个时间，分别测试一下是否为之前、之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 `Sub` 返回一个 `Duration` 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	// 我们可以用各种单位来表示时间段的长度。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用 `Add` 将时间后移一个时间段，或者使用一个 `-` 来将时间前移一个时间段。
	p(then.Add(diff))
	p(then.Add(-diff))
}

func TimeFormatExample() {
	p := fmt.Println

	// 这是一个遵循 RFC3339，
	// 并使用对应的 `布局`（layout）常量进行格式化的基本例子。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 时间解析使用与 `Format` 相同的布局值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` 和 `Parse` 使用基于例子的布局来决定日期格式，
	// 一般你只要使用 `time` 包中提供的布局常量就行了，但是你也可以实现自定义布局。
	// 布局时间必须使用 `Mon Jan 2 15:04:05 MST 2006` 的格式，
	// 来指定 格式化/解析给定时间/字符串 的布局。
	// 时间一定要遵循：2006 为年，15 为小时，Monday 代表星期几等规则。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示的时间（时间戳），
	// 您还可以将标准字符串格式与提取时间值的一部分一起使用。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// 当输入的时间格式不正确时，`Parse` 会返回一个解析错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
