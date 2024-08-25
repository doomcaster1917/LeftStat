package charts_img_maker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"sort"
)

type ChartImager interface {
	makeChartImg() *os.File
}

type Chart struct {
	XValues []float64
	YValues []float64
}

func (c Chart) MakeChartImg() string {
	sort.Float64s(c.YValues)
	graph := chart.Chart{
		YAxis: chart.YAxis{
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%.0f", vf)
				}
				return ""
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: []float64{2012.00, 2013.00, 2014.00, 2015.00, 2016.00, 2017.00, 2018.00, 2019.00, 2020.00, 2021.00, 2022.00, 2023.00}, //c.XValues,

				YValues: c.YValues,
			},
		},
	}

	id := uuid.New()
	name := id.String()
	f, _ := os.Create(fmt.Sprintf("static/%v.png", name))
	defer f.Close()
	graph.Render(chart.PNG, f)

	return name + ".png"
}
