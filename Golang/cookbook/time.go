package cookbook

import "time"

// 1659977921597
func Unix2Time(timestamp int64) string {
	return time.UnixMilli(timestamp).In(time.FixedZone("UTC+8", 8*3600)).Format("2006-01-02 15:04:05")
}
