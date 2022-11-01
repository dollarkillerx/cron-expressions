package cron_expressions

import "strconv"

/**

每天幾點: 0 0 19 * * ?         // 每天19 點

每周几: 0 0 0 ? * 1      // 每周1     0點

每月： 0 0 0 2 * ?      // 每月 的 2 號

second, minute, hour, day, moon, week
*/

type CronExpressions struct {
	second string
	minute string
	hour   string
	day    string
	moon   string
	week   string
}

func Expressions() *CronExpressions {
	return &CronExpressions{
		second: "0",
		minute: "0",
		hour:   "0",
		day:    "?",
		moon:   "*",
		week:   "?", // 不支持同时指定星期几和几月参数。 Specifying both the day of the week and the month parameter at the same time is not supported.
	}

	// '？'只能在日和星期（周）中指定使用，其作用为不指定
	// '? 'Can only be specified in day and week (week), its effect is not specified
}

// Second 秒
func (c *CronExpressions) Second(second uint) *CronExpressions {
	if second > 60 {
		panic("CronExpressions Second > 60")
	}

	c.second = strconv.Itoa(int(second))

	return c
}

// EverySecond 每秒
func (c *CronExpressions) EverySecond() *CronExpressions {
	c.second = "*"
	return c
}

// Minute  分鐘
func (c *CronExpressions) Minute(minute uint) *CronExpressions {
	if minute > 60 {
		panic("CronExpressions Minute > 60")
	}

	c.minute = strconv.Itoa(int(minute))

	return c
}

// EveryMinute 每分鐘
func (c *CronExpressions) EveryMinute() *CronExpressions {
	c.minute = "*"
	return c
}

// Hour  小時
func (c *CronExpressions) Hour(hour uint) *CronExpressions {
	if hour > 24 {
		panic("CronExpressions Hour > 24")
	}

	c.hour = strconv.Itoa(int(hour))

	return c
}

// EveryHour 每小時
func (c *CronExpressions) EveryHour() *CronExpressions {
	c.hour = "*"
	return c
}

// Day  天
func (c *CronExpressions) Day(day uint) *CronExpressions {
	if day > 31 {
		panic("CronExpressions Day > 31")
	}

	c.day = strconv.Itoa(int(day))

	return c
}

// EveryDay 每天
func (c *CronExpressions) EveryDay() *CronExpressions {
	c.day = "*"
	return c
}

// Moon 月
func (c *CronExpressions) Moon(moon uint) *CronExpressions {
	if moon > 7 {
		panic("CronExpressions Moon > 7")
	}

	c.moon = strconv.Itoa(int(moon))

	return c
}

// EveryMoon 每月
func (c *CronExpressions) EveryMoon() *CronExpressions {
	c.moon = "*"
	return c
}

// Week 周
func (c *CronExpressions) Week(week uint) *CronExpressions {
	if week > 7 {
		panic("CronExpressions Week > 7")
	}

	c.week = strconv.Itoa(int(week))

	return c
}

// EveryWeek 每周
func (c *CronExpressions) EveryWeek() *CronExpressions {
	c.week = "*"
	return c
}
