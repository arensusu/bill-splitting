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
	<div class="container-md">
		<span class="px-5">Filters:</span>
		<div class="btn-group" role="group" aria-label="button group">
			{{range (index .MultiSeries 0).Data}}
				<button type="button" class="btn" data-bs-toggle="button">{{.Name}}</button>
			{{end}}
		</div>
	</div>
	<div class="container-md overflow-y-auto" style="max-height: 600px">
		<table class="table" id="sortTable">
			<thead>
				<tr>
					<th class="position-relative" onclick="sortTable(0)">Date <span class="position-absolute end-0">&#x25b4;&#x25be;</span></th>
					<th class="position-relative" onclick="sortTable(1)">Category <span class="position-absolute end-0">&#x25b4;&#x25be;</span></th>
					<th class="position-relative" onclick="sortTable(2)">Description <span class="position-absolute end-0">&#x25b4;&#x25be;</span></th>
					<th class="position-relative" onclick="sortTable(3)">Price <span class="position-absolute end-0">&#x25b4;&#x25be;</span></th>
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
	body {font-family: monospace;}
</style>
<script
  src="https://code.jquery.com/jquery-3.7.1.min.js"
  integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo="
  crossorigin="anonymous">
</script>
<script>

$(document).ready(function () {
	$("button").click(function () {
		var classes = [];

		$("button").each(function () {
			if ($(this).is(".active")) { classes.push($(this).text()); }
		});

		if (classes.length == $("button").length) {
			$("button").each(function () {
				$(this).removeClass("active");
			});

			classes = [];
		}

		if (classes.length == 0) { // if no filters selected, show all items
			$("#sortTable tbody tr").show();
		} else { // otherwise, hide everything...
			$("#sortTable tbody tr").hide();

			$("#sortTable tr").each(function () {
				var show = false;
				var row = $(this);
				classes.forEach(function (className) {
					if (row.find('td').eq(1).text() == className) { show = true; }
				});
				if (show) { row.show(); }
			});
		}
	});
});

function sortTable(n) {
  var table, rows, switching, i, x, y, shouldSwitch, dir, switchcount = 0;
  table = document.getElementById("sortTable");
  switching = true;
  // Set the sorting direction to ascending:
  dir = "asc";
  /* Make a loop that will continue until
  no switching has been done: */
  while (switching) {
    // Start by saying: no switching is done:
    switching = false;
    rows = table.rows;
    /* Loop through all table rows (except the
    first, which contains table headers): */
    for (i = 1; i < (rows.length - 1); i++) {
      // Start by saying there should be no switching:
      shouldSwitch = false;
      /* Get the two elements you want to compare,
      one from current row and one from the next: */
      x = rows[i].getElementsByTagName("TD")[n];
      y = rows[i + 1].getElementsByTagName("TD")[n];
      /* Check if the two rows should switch place,
      based on the direction, asc or desc: */
      if (dir == "asc") {
        if (x.innerHTML.toLowerCase() > y.innerHTML.toLowerCase()) {
          // If so, mark as a switch and break the loop:
          shouldSwitch = true;
          break;
        }
      } else if (dir == "desc") {
        if (x.innerHTML.toLowerCase() < y.innerHTML.toLowerCase()) {
          // If so, mark as a switch and break the loop:
          shouldSwitch = true;
          break;
        }
      }
    }
    if (shouldSwitch) {
      /* If a switch has been marked, make the switch
      and mark that a switch has been done: */
      rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
      switching = true;
      // Each time a switch is done, increase this count by 1:
      switchcount ++;
    } else {
      /* If no switching has been done AND the direction is "asc",
      set the direction to "desc" and run the while loop again. */
      if (switchcount == 0 && dir == "asc") {
        dir = "desc";
        switching = true;
      }
    }
  }
}
</script>
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
			TitleStyle:    &opts.TextStyle{Color: "black", FontSize: 28, FontFamily: "monospace"},
			Subtitle:      subtitle,
			SubtitleStyle: &opts.TextStyle{Color: "black", FontSize: 20, FontFamily: "monospace"},
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
		data = append(data, opts.PieData{Value: fmt.Sprintf("%.0f", values[i]/total*100), Name: legends[i]})
	}

	for _, dataset := range datasets {
		pie.AddDataset(opts.Dataset{Source: dataset})
	}

	pie.AddSeries("", data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:       opts.Bool(true),
				Formatter:  "{b}: {c}%",
				FontSize:   16,
				FontFamily: "monospace",
			}),
		)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", pie)

	return pie.Render(io.MultiWriter(f))
}
