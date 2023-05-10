package main

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"gostudy/libs"
	"math/big"
	"net"
	"reflect"
	"strconv"
	"time"
)

func main() {
	user := &study.User{Name: "赵钱孙李", Age: 32}
	fmt.Println(user.MyAge())
	var a study.Action = user
	a.SetAge(45)
	fmt.Println(user.MyAge())
}

func studyConcat() {
	a := "sss"
	b := "ddd"
	fmt.Println(a + b)
}

func studyIpNetwork() {
	net.ParseIP("1.2.3.4")

}

func studyStr2num() {
	ia := 10
	var iaaa int64
	ib := "50000000000000000000"

	fa := 100.4
	fb := "50.6"

	iaa := strconv.Itoa(ia)
	ibb, _ := strconv.Atoi(ib)
	ibb2, _ := strconv.ParseInt(ib, 10, 64)

	faa := strconv.FormatFloat(fa, 'f', -1, 32)
	fbb, _ := strconv.ParseFloat(fb, 32)
	fbb2 := float32(fbb)

	fmt.Println(new(big.Int).SetString(ib, 10))
	fmt.Println(iaa, reflect.TypeOf(iaa))
	fmt.Println(ibb, reflect.TypeOf(ibb))
	fmt.Println(ibb2, reflect.TypeOf(ibb2))
	fmt.Println(faa, reflect.TypeOf(faa))
	fmt.Println(fbb, reflect.TypeOf(fbb))
	fmt.Println(fbb2, reflect.TypeOf(fbb2))
	fmt.Println(iaaa)
}

func studyDatetime() {
	now := time.Now()
	utc := time.Now().UTC()
	tz, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().Unix())
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(utc.In(tz).Add(-2 * time.Hour))
	fmt.Println(time.Unix(1683275206, 0))
	var dt1 = time.Date(2013, 12, 20, 0, 0, 0, 0, tz)
	var dt2 = time.Date(2014, 1, 20, 0, 0, 0, 0, tz)
	for dt := dt1; dt.Before(dt2); dt = dt.Add(time.Hour * 24) {
		fmt.Println(dt)
	}

	var nowc = carbon.FromStdTime(now)
	// 时区转换
	fmt.Println(nowc.SetTimezone("utc"), "xxxxxxxx")
	// 修改日期
	fmt.Println(nowc.SetDay(25), "xxxxxxxx")
	fmt.Println(nowc.SubDays(30).Date())
	fmt.Println(carbon.CreateFromTimestamp(1683275206).ToDateTimeString())
	var dtc1 = carbon.CreateFromDate(2021, 12, 20)
	for dt := dtc1; dt.Lte(dtc1.AddDays(30)); dt = dt.AddDay() {
		fmt.Println(dt.ToIso8601String())
		fmt.Println(dt.Format("Y-m-d H:i:s P"))
	}
}
