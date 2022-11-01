# Cron Expressions


``` 
	//  每天 19 點, every day at 19:00

	v1 := Expressions().EveryDay().Hour(19)

	fmt.Println("每天19 點: ", v1)
	fmt.Println("every day at 19:00: ", v1)

	//  每周1, 0點,   monday, 0:00

	v2 := Expressions().Week(1).Hour(0)

	fmt.Println("每周1, 0點: ", v2)
	fmt.Println("monday, 0:00: ", v2)

	//  每月 的 2 號,  2nd of the month

	v3 := Expressions().Day(2)

	fmt.Println("每月 的 2 號: ", v3)
	fmt.Println("2nd of the month: ", v3)

```

Demo: 

``` 
	//  每天 19 點

	v1 := Expressions().EveryDay().Hour(19)

	fmt.Println("每天19 點: ", v1.String(), v1.NextTime(), v1.StringAccurate())

	//  每周1, 0點

	v2 := Expressions().Week(1).Hour(0)

	fmt.Println("每周1, 0點: ", v2.String(), v2.NextTime(), v2.StringAccurate())

	//  每月 的 2 號

	v3 := Expressions().Day(3)

	fmt.Println("每月 的 2 號: ", v3.String(), v3.NextTime(), v3.StringAccurate())
	
	每天19 點:  0 19 * * ? 2022-11-01 19:00:00 +0800 CST 0 0 19 * * ?
    每周1, 0點:  0 0 ? * 1 2022-11-07 00:00:00 +0800 CST 0 0 0 ? * 1
    每月 的 2 號:  0 0 3 * ? 2022-11-03 00:00:00 +0800 CST 0 0 0 3 * ?
```