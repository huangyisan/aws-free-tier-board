package common

import (
	"time"
)

var (
	AllTrafficQuery = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`
)

func GetMonthFirstDay() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	// 2019-08-28T22:00:00Z
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	//return firstOfMonth.Format("2006-01-02T15:04:05Z")
	return firstOfMonth.Format(time.RFC3339)

}

func GetCurrentDay() time.Time {
	now := time.Now()
	return now
}
