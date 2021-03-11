package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)                                              //获取当前时间
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) //格式化当前时间

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) //2021-03-17 19:20:47.9640888 +0800 CST m=+604800.002206401

	fmt.Println(t.Format(time.RFC822)) //10 Mar 21 19:20 CST

	fmt.Println(t.Format(time.ANSIC)) //Wed Mar 10 19:21:23 2021

	fmt.Println(t.Format("21 Dec 2011 08:52"))

	s := t.Format("20060102")
	fmt.Println(t, "=>", s) // 2021-03-10 19:23:40.5056711 +0800 CST m=+0.001609101 => 20210310
}
