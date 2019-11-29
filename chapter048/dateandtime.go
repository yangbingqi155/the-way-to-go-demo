package main

import (
	"fmt"
	"time"
)

type cnTime struct {
	time.Time
}

func (t *cnTime) String() (cnDateFormat string) {
	return fmt.Sprintf("%d年%d月%d日", t.Year(), t.Month(), t.Day())
}

func main() {
	fmt.Println("Date and time")
	now := cnTime{time.Now()}
	fmt.Printf("Current Time:%s\n", now.String())
	fmt.Printf("Current Time(custom format):%s\n", now.Format("20190101 16:00:00"))
	fmt.Printf("Current Time(UTC):%v", now.UTC())
}
