package main

import (
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart
func idealDuration() []opts.LineData {
	items := make([]opts.LineData, 0)
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 24})
	items = append(items, opts.LineData{Value: 18})
	items = append(items, opts.LineData{Value: 12})
	items = append(items, opts.LineData{Value: 6})
	items = append(items, opts.LineData{Value: 0})

	return items
}
func actualDuration() []opts.LineData {
	items := make([]opts.LineData, 0)
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 25})
	items = append(items, opts.LineData{Value: 19.5})
	items = append(items, opts.LineData{Value: 12})
	items = append(items, opts.LineData{Value: 7.5})
	items = append(items, opts.LineData{Value: 0})

	return items
}
func totalProjectHours() []opts.LineData {
	items := make([]opts.LineData, 0)
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30})

	return items
}
func remainingEffortPlusLogHours() []opts.LineData {
	items := make([]opts.LineData, 0)
	items = append(items, opts.LineData{Value: 30})
	items = append(items, opts.LineData{Value: 30.25})
	items = append(items, opts.LineData{Value: 30.25})
	items = append(items, opts.LineData{Value: 31.25})
	items = append(items, opts.LineData{Value: 32.25})
	items = append(items, opts.LineData{Value: 31.25})

	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Weekly Sprint"}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Weekly Sprint",
			Subtitle: "",
		}),
	)
	// Put data into instance
	line.SetXAxis([]string{"Plan", "Mon", "Tue", "Wed", "Thu", "Fri"}).
		AddSeries("Ideal Duration", idealDuration(), charts.WithLineStyleOpts(opts.LineStyle{Color: "green", Type: "dashed"})).
		AddSeries("Actual Duration", actualDuration(), charts.WithLineStyleOpts(opts.LineStyle{Color: "orange", Type: "solid"})).
		AddSeries("Total Project Hours", totalProjectHours(), charts.WithLineStyleOpts(opts.LineStyle{Color: "green", Type: "dashed"})).
		AddSeries("Remaining Effort Plus Log Hours", remainingEffortPlusLogHours(), charts.WithLineStyleOpts(opts.LineStyle{Color: "purple", Type: "solid"})).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
