package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/templates"
)

var chartTml = `{{- define "chart" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
    {{- template "base" . }}
	<div class="container">
		<table class="table table-striped">
			<thead>
				<tr>
					<th>Date</th>
					<th>Category</th>
					<th>Description</th>
					<th>Price</th>
				</tr>
			</thead>
			<tbody>
				{{range .DatasetList}}
					<tr>
					{{range .Source}}
						<td>{{.}}</td>
					{{end}}
					</tr>
				{{end}}
			</tbody>
		</table>
	</div>
<style>
    .container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>
</body>
</html>
{{ end }}
`

type customRender struct {
	render.BaseRender
	c      interface{}
	before []func()
}

func NewCustomRender(c interface{}, before ...func()) render.Renderer {
	return &customRender{c: c, before: before}
}

func (r *customRender) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	r.before = []func(){}

	contents := []string{templates.HeaderTpl, templates.BaseTpl, chartTml}
	tpl := render.MustTemplate(render.ModChart, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, render.ModChart, r.c); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}

func CreatePieChart(values []float64, legends []string, title, subtitle string, datasets [][4]string, path string) error {
	pie := charts.NewPie()
	pie.Renderer = NewCustomRender(pie, pie.Validate)

	pie.SetGlobalOptions(
		charts.WithAnimation(false),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "white",
			PageTitle:       "Expenses",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:         title,
			TitleStyle:    &opts.TextStyle{Color: "black", FontSize: 28},
			Subtitle:      subtitle,
			SubtitleStyle: &opts.TextStyle{Color: "black", FontSize: 20},
			Left:          "80",
			Top:           "20",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   opts.Bool(false),
			Left:   "20",
			Top:    "20",
			Orient: "vertical",
		}),
	)

	pie.AddCustomizedJSAssets("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js")
	pie.AddCustomizedCSSAssets("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css")

	total := 0.0
	for _, value := range values {
		total += value
	}

	data := make([]opts.PieData, 0)
	for i := 0; i < len(values); i++ {
		data = append(data, opts.PieData{Value: values[i], Name: legends[i]})
	}

	for _, dataset := range datasets {
		pie.AddDataset(opts.Dataset{Source: dataset})
	}

	pie.AddSeries("pie", data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}%",
				FontSize:  16,
			}),
		)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", pie)

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
