package encharts_maker

import (
	"bytes"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	chartrender "github.com/go-echarts/go-echarts/v2/render"
	"html/template"
	"log"
	"sort"
)

type DataSets struct {
	Data map[int]int
	Id   int
	Name string
}

type Chart struct {
	Id          int        `json:"Id"`
	Order       int        `json:"Order"`
	Name        string     `json:"Name"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DataSets    []DataSets `json:"DataSets"`
	MainAxisId  int        `json:"MainAxisId"`
}

var test = map[int]int{2013: 2758, 2014: 2996, 2015: 3516, 2016: 3632, 2017: 3729, 2018: 3840, 2019: 4062, 2020: 4267, 2021: 4890, 2022: 5500, 2023: 5722}

func deserializeY(data map[int]int) []opts.LineData {
	yPoints := make([]opts.LineData, 0)
	keys := make([]int, 0)
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		yPoints = append(yPoints, opts.LineData{Value: data[k]})
	}

	return yPoints
}

func deserializeX(data map[int]int) []int {
	yPoints := make([]int, 0)
	keys := make([]int, 0)
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		yPoints = append(yPoints, k)
	}

	return yPoints
}

func GenerateChart(chart Chart) template.HTML {

	bar := charts.NewLine()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    chart.Name,
		Subtitle: chart.Title,
	}))

	x := make([]int, 0)

	for _, val := range chart.DataSets {
		if chart.MainAxisId == val.Id {
			x = deserializeX(val.Data)
			y := deserializeY(val.Data)
			bar.SetXAxis(x).AddSeries(val.Name, y)
		} else {
			y := deserializeY(val.Data)
			bar.AddSeries(val.Name, y)
		}
	}

	var htmlSnippet template.HTML = renderToHtml(bar)
	return htmlSnippet
}

func renderToHtml(c interface{}) template.HTML {
	var buf bytes.Buffer
	r := c.(chartrender.Renderer)
	err := r.Render(&buf)
	if err != nil {
		log.Printf("Failed to render chart: %s", err)
		return ""
	}

	return template.HTML(buf.String())
}
