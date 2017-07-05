package main

import (
	"fmt"
	"time"

	now "github.com/jinzhu/now"
)

func main() {
	p := fmt.Println
	p(time.Now()) // 2013-11-18 17:51:49.123456789 Mon

	p(now.BeginningOfHour) // 2013-11-18 17:51:00 Mon
	// now.BeginningOfHour()     // 2013-11-18 17:00:00 Mon
	// now.BeginningOfDay()      // 2013-11-18 00:00:00 Mon
	// now.BeginningOfWeek()     // 2013-11-17 00:00:00 Sun
	// now.FirstDayMonday = true // Set Monday as first day, default is Sunday
	// now.BeginningOfWeek()     // 2013-11-18 00:00:00 Mon
	// now.BeginningOfMonth()    // 2013-11-01 00:00:00 Fri
	// now.BeginningOfQuarter()  // 2013-10-01 00:00:00 Tue
	// now.BeginningOfYear()     // 2013-01-01 00:00:00 Tue

	// now.EndOfMinute()         // 2013-11-18 17:51:59.999999999 Mon
	// now.EndOfHour()           // 2013-11-18 17:59:59.999999999 Mon
	p(now.EndOfDay()) // 2013-11-18 23:59:59.999999999 Mon
	// now.EndOfWeek()           // 2013-11-23 23:59:59.999999999 Sat
	// now.FirstDayMonday = true // Set Monday as first day, default is Sunday
	p(now.EndOfWeek()) // 2013-11-24 23:59:59.999999999 Sun
	// now.EndOfMonth()          // 2013-11-30 23:59:59.999999999 Sat
	// now.EndOfQuarter()        // 2013-12-31 23:59:59.999999999 Tue
	p(now.EndOfYear()) // 2013-12-31 23:59:59.999999999 Tue
}
