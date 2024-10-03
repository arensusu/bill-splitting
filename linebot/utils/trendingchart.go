package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func CreateTrendingChart(values []map[string]float64, xAxis []string, title, path string) error {
	bar := charts.NewBar()
	//pie.Renderer = NewCustomRender(pie, pie.Validate)

	bar.SetGlobalOptions(
		charts.WithAnimation(false),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "white",
			PageTitle:       "Expenses",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:      title,
			TitleStyle: &opts.TextStyle{Color: "black", FontSize: 28, FontFamily: "monospace"},
			Left:       "80",
			Top:        "20",
		}),
	)

	dataNum := len(values)
	data := map[string][]opts.BarData{}
	lineData := make([]opts.LineData, dataNum)

	for i := 0; i < len(values); i++ {
		total := 0.0
		for category, value := range values[i] {
			if _, ok := data[category]; !ok {
				data[category] = make([]opts.BarData, dataNum)
			}

			data[category][i] = opts.BarData{Value: value, Name: category}
			total += value
		}
		lineData[i] = opts.LineData{Value: total}
	}

	line := charts.NewLine()
	line.AddSeries("", lineData)

	for _, v := range data {
		bar.AddSeries("", v)
	}

	bar.SetXAxis(xAxis).Overlap(line)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", bar)

	return bar.Render(io.MultiWriter(f))
}
