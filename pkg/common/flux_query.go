package common

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
		|> map(fn: (r) => ({ r with download: math.abs(x: r.download), upload: math.abs(x: r.upload)}))`

	AllTrafficByGroup = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")
			
		download_sum = union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with download: math.abs(x: r.download), upload: math.abs(x: r.upload)}))
		|> group(columns: ["group"])
        |> sum(column: "download")
        |> group()

		upload_sum = union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with download: math.abs(x: r.download), upload: math.abs(x: r.upload)}))
		|> group(columns: ["group"])
        |> sum(column: "upload")
        |> group()
		join(tables: {t1: download_sum, t2: upload_sum}, on: ["group"])`

	TrafficByTag = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["tag"] == "%[4]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan")
			|> filter(fn: (r) => r["tag"] == "%[4]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`

	TrafficByGroup = `import "math"
		last_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["group"] == "%[4]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "%[1]v") 
			|> range(start: %[2]v, stop: %[3]v)
			|> filter(fn: (r) => r["_measurement"] == "trojan")
			|> filter(fn: (r) => r["group"] == "%[4]v")
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")
		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
        |> map(fn: (r) => ({ r with download: math.abs(x: r.download), upload: math.abs(x: r.upload)}))`
)
