package cron_expressions

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aptible/supercronic/cronexpr"
	"github.com/robfig/cron"
)

func TestCron1(t *testing.T) {
	//  每天 19 點

	v1 := Expressions().EveryDay().Hour(19)

	fmt.Println("每天19 點: ", v1.String(), v1.NextTime(), v1.StringAccurate())

	//  每周1, 0點

	v2 := Expressions().Week(1).Hour(0)

	fmt.Println("每周1, 0點: ", v2.String(), v2.NextTime(), v2.StringAccurate())

	//  每月 的 2 號

	v3 := Expressions().Day(3)

	fmt.Println("每月 的 2 號: ", v3.String(), v3.NextTime(), v3.StringAccurate())
}

func TestCron(t *testing.T) {
	cron := cron.New()
	err := cron.AddFunc("0 0 2 * ?", func() {
		fmt.Println("aaaa")
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func TestCp2(t *testing.T) {
	v2 := Expressions().Day(2)

	next := cronexpr.MustParse(v2.String()).Next(time.Now()) // 這個不能精確的秒
	next2 := cronexpr.MustParse(v2.String()).Next(next)      // 這個不能精確的秒
	fmt.Println(next.String())
	fmt.Println(next2.String())
}
