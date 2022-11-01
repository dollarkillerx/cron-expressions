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