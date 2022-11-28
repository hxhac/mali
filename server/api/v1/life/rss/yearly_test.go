package rss

import (
	htime "github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"github.com/golang-module/carbon"
	"testing"
)

const (
	TwoDaily         = "@2daily"
	ThreeDaily       = "@3daily"
	FourDaily        = "@4daily"
	SixDaily         = "@6daily"
	Weekly           = "@weekly"
	TwoWeekly        = "@2weekly"
	FourWeekly       = "@4weekly"
	EightWeekly      = "@8weekly"
	TwentyFourWeekly = "@24weekly"
	Monthly          = "@monthly"
	TwoMonthly       = "@2monthly"
	ThreeMonthly     = "@3monthly"
	SixMonthly       = "@6monthly"
	Yearly           = "@yearly"
)

// TODO weekly的第二天有bug，但是暂时不用处理
func TestCheckCron(t *testing.T) {
	type args struct {
		cb        carbon.Carbon
		cronTime  string
		isWeekDay bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-03")}, true},
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-04")}, false},
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-05")}, true},
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-06")}, false},
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-30")}, false},
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-12-31")}, true}, // 转过年有问题
		// {TwoDaily, args{cronTime: TwoDaily, cb: xzb("2021-1-02")}, false},

		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-03")}, true},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-04")}, false},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-05")}, false},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-06")}, true},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-30")}, true},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-12-31")}, false}, // 转过年有问题
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2022-1-01")}, true},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-1-02")}, false},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-1-03")}, false},
		// {ThreeDaily, args{cronTime: ThreeDaily, cb: xzb("2021-1-04")}, true},

		{Weekly, args{xzb("2021-12-04"), Weekly, true}, true},
		{Weekly, args{xzb("2021-12-11"), Weekly, true}, true},          // wn=0
		{Weekly + "第二天", args{xzb("2021-12-12"), Weekly, true}, false}, // wn=0
		{Weekly, args{xzb("2021-12-18"), Weekly, true}, true},
		{Weekly, args{xzb("2022-03-12"), Weekly, true}, true},
		{Weekly, args{xzb("2022-03-19"), Weekly, true}, true},
		{Weekly, args{xzb("2022-03-26"), Weekly, true}, true},
		{Weekly + "第二天", args{xzb("2022-03-27"), Weekly, true}, false}, //
		{Weekly, args{xzb("2022-11-05"), Weekly, true}, true},
		{Weekly, args{xzb("2022-11-12"), Weekly, true}, true},
		{Weekly, args{xzb("2022-11-19"), Weekly, true}, true},
		{Weekly, args{xzb("2022-11-26"), Weekly, true}, true},
		{Weekly, args{xzb("2022-12-03"), Weekly, true}, true},

		// @2weekly
		{TwoWeekly, args{xzb("2022-11-05"), TwoWeekly, true}, false},
		{TwoWeekly, args{xzb("2022-11-12"), TwoWeekly, true}, true},
		{TwoWeekly, args{xzb("2022-11-19"), TwoWeekly, true}, false},
		{TwoWeekly, args{xzb("2022-11-26"), TwoWeekly, true}, true},
		{TwoWeekly, args{xzb("2022-12-03"), TwoWeekly, true}, false},
		{TwoWeekly + "第二天", args{xzb("2021-12-04"), TwoWeekly, true}, false}, // 第二天
		{TwoWeekly, args{xzb("2021-12-11"), TwoWeekly, true}, true},
		{TwoWeekly + "第二天", args{xzb("2021-12-12"), TwoWeekly, true}, false}, // 第二天
		{TwoWeekly, args{xzb("2021-12-18"), TwoWeekly, true}, false},
		{TwoWeekly, args{xzb("2021-12-25"), TwoWeekly, true}, true},
		{TwoWeekly, args{xzb("2022-01-08"), TwoWeekly, true}, true},
		{TwoWeekly, args{xzb("2022-01-15"), TwoWeekly, true}, false},
		{TwoWeekly, args{xzb("2022-01-22"), TwoWeekly, true}, true},

		// @4weekly
		{FourWeekly, args{xzb("2022-11-05"), FourWeekly, true}, false},
		{FourWeekly, args{xzb("2022-11-12"), FourWeekly, true}, true},
		{FourWeekly, args{xzb("2022-11-19"), FourWeekly, true}, false}, // wn=1
		{FourWeekly, args{xzb("2022-11-26"), FourWeekly, true}, false}, // wn=2
		{FourWeekly, args{xzb("2022-12-03"), FourWeekly, true}, false},
		{FourWeekly, args{xzb("2022-12-10"), FourWeekly, true}, true},
		{FourWeekly, args{xzb("2022-12-17"), FourWeekly, true}, false},

		// @8weekly
		{EightWeekly, args{xzb("2022-10-08"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-10-15"), EightWeekly, true}, true},
		{EightWeekly, args{xzb("2022-10-22"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-10-29"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-11-05"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-11-12"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-11-19"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-11-26"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-12-03"), EightWeekly, true}, false},
		{EightWeekly, args{xzb("2022-12-10"), EightWeekly, true}, true},
		{EightWeekly, args{xzb("2022-12-17"), EightWeekly, true}, false},
		

		// {Monthly, args{cronTime: Monthly, cb: xzb("2021-12-01")}, true},
		// {Monthly, args{cronTime: Monthly, cb: xzb("2022-4-01")}, true},
		// {Monthly, args{cronTime: Monthly, cb: xzb("2021-12-02")}, false},
		//
		// {TwoMonthly, args{cronTime: TwoMonthly, cb: xzb("2021-3-01")}, true},
		// {TwoMonthly, args{cronTime: TwoMonthly, cb: xzb("2021-5-01")}, true},
		// {TwoMonthly, args{cronTime: TwoMonthly, cb: xzb("2021-6-01")}, false},
		//
		// {ThreeMonthly, args{cronTime: ThreeMonthly, cb: xzb("2021-4-01")}, true},
		// {ThreeMonthly, args{cronTime: ThreeMonthly, cb: xzb("2021-10-01")}, true},
		// {ThreeMonthly, args{cronTime: ThreeMonthly, cb: xzb("2021-11-01")}, false},
		//
		// {SixMonthly, args{cronTime: SixMonthly, cb: xzb("2021-7-01")}, true},
		// {SixMonthly, args{cronTime: SixMonthly, cb: xzb("2021-8-01")}, false},
		//
		// {Yearly, args{cronTime: Yearly, cb: xzb("2022-1-01")}, true},
		// {Yearly, args{cronTime: Yearly, cb: xzb("2022-1-02")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCron(tt.args.cb, tt.args.cronTime, tt.args.isWeekDay); got != tt.want {
				t.Errorf("CheckCron() = %v, want %v", got, tt.want)
			}
		})
	}
}

func xzb(tt string) carbon.Carbon {
	zx := htime.StrToTime(tt, "Y-m-d")
	return carbon.Time2Carbon(zx)
}
