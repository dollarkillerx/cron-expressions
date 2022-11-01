package cron_expressions

import (
	"fmt"
	"log"
	"testing"

	"github.com/robfig/cron"
)

func TestCron1(t *testing.T) {
	//  每天 19 點

	v1 := Expressions().EveryDay().Hour(19)

	fmt.Println("每天19 點: ", v1)

	//  每周1, 0點

	v2 := Expressions().Week(1).Hour(0)

	fmt.Println("每周1, 0點: ", v2)

	//  每月 的 2 號

	v3 := Expressions().Day(2)

	fmt.Println("每月 的 2 號: ", v3)
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
