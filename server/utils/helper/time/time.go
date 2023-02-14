package time

import (
	"fmt"
	"time"

	"github.com/gogf/gf/os/gtime"
)

// GetToday 获取今天的零点时间
func GetToday() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}

// StrToTime 字符串转time.Time
func StrToTime(str, format string) time.Time {
	st, err := gtime.StrToTimeFormat(str, format)
	if err != nil {
		return time.Time{}
	}
	return st.Time
}

func TransTime(str string) time.Time {
	format, err := gtime.StrToTimeFormat(str, "Y/n/d H:i:s")
	if err != nil {
		return time.Time{}
	}
	return format.Time
}

// 与当前时间对比
func CheckDateTime(nn string) *gtime.Time {
	str, err := gtime.NewFromTime(GetToday()).AddStr(nn)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return str
}

func DiffDateTime(nn string) int64 {
	str, _ := gtime.NewFromTime(GetToday()).AddStr(nn)
	seconds := str.Sub(gtime.NewFromTime(GetToday())).Milliseconds()
	return seconds
}

func DiffNowDateTime(nn string) int64 {
	str, _ := gtime.NewFromTime(GetToday()).AddStr(nn)
	seconds := gtime.Now().Sub(str).Milliseconds()
	return seconds
}

// CheckTimeLimit 判断时间范围
func CheckTimeLimit(nn string) bool {
	diff := DiffNowDateTime(nn)
	// 没到目标时间
	if diff < 0 {
		return false
	}

	// 超过目标时间30min
	duration, err := time.ParseDuration("30m")
	if err != nil {
		return false
	}
	ss := duration.Milliseconds()
	if diff > ss {
		return false
	}

	return true
}
