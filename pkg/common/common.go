package common

import (
	"time"
)

const (
	layoutISO = "2006-01-02"
)

var (
	AllTraffic = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`

	AllTrafficByGroup = `import "math"
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
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))
        |> group(columns: ["group"])
        |> sum(column: "total")
        |> group()`

	TrafficByTag = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["tag"] == "%[3]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan")
			|> filter(fn: (r) => r["tag"] == "%[3]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`

	TrafficByGroup = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["group"] == "%[3]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan")
			|> filter(fn: (r) => r["group"] == "%[3]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`
)

func FormatDateToRFC3339(date string) string {
	t, _ := time.Parse(layoutISO, date)
	nt := t.Format(time.RFC3339)
	return nt
}

func GetMonthFirstDay() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	// 2019-08-28T22:00:00Z
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return firstOfMonth.Format(time.RFC3339)

}

func GetCurrentDay() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
