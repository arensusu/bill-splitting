package utils

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func CreatePieChart(values []any, legends []string, title, subtitle string, path string) error {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithAnimation(false),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "white",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: subtitle,
			Left:     "20",
			Top:      "20",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   opts.Bool(false),
			Left:   "20",
			Top:    "20",
			Orient: "vertical",
		}),
	)

	data := make([]opts.PieData, 0)
	for i := 0; i < len(values); i++ {
		data = append(data, opts.PieData{Value: values[i], Name: legends[i]})
	}

	pie.AddSeries("pie", data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
		)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return pie.Render(io.MultiWriter(f))
	// return render.MakeChartSnapshot(pie.RenderContent(), path)

	// fontBytes, err := os.ReadFile("../msjh.ttc")
	// if err != nil {
	// 	return err
	// }

	// font, err := truetype.Parse(fontBytes)
	// if err != nil {
	// 	return err
	// }

}
