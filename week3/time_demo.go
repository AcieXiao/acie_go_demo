package main

import (
	"fmt"
	"time"
)

func main() {
	unix := time.Now().Unix()

	fmt.Println("unix = ", unix)
	fmt.Println("UnixMilli = ", time.Now().UnixMilli())

	nano := time.Now().UnixNano()
	fmt.Printf("当前时间戳(单位纳秒): %v\n", nano)

	format := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("当前时间(Y-m-d H:i:s): %v\n", format)

	format2 := time.Now().Format("20060102150405")
	fmt.Printf("当前时间(YmdHis): %v\n", format2)

	fmt.Printf("当前年: %v\n", time.Now().Year())
	fmt.Printf("当前月: %v\n", time.Now().Month())
	fmt.Printf("当前日: %v\n", time.Now().Day())
	fmt.Printf("当前小时: %v\n", time.Now().Hour())
	fmt.Printf("当前分钟: %v\n", time.Now().Minute())
	fmt.Printf("当前秒: %v\n", time.Now().Second())
	fmt.Printf("当前星期几: %v\n", time.Now().Weekday())

	// 十分钟前
	now := time.Now()
	duration, _ := time.ParseDuration("-10m")
	fmt.Printf("十分钟前: %v \n", now.Add(duration))
	// 一小时前
	duration2, _ := time.ParseDuration("-1h")
	fmt.Printf("一小时前: %v \n", now.Add(duration2))
	// 一天后
	fmt.Printf("一天后: %v \n", now.AddDate(0, 0, 1))
	fmt.Printf("一月后: %v \n", now.AddDate(0, 1, 0))
	fmt.Printf("一年后: %v \n", now.AddDate(1, 0, 0))

}
