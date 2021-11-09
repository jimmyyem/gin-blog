package util

import "time"

// 根据时间秒数转固定格式
func FormatAsTime(second int) string {
	return time.Unix(int64(second), 0).Format("2006-01-02 15:04:05")
}