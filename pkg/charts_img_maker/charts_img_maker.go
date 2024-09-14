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
	sort.Float64s(c.XValues)

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
				XValues: c.XValues,
				YValues: c.YValues,
			},
		},
	}

	id := uuid.New()
	name := id.String()
	f, err := os.Create(fmt.Sprintf("static/%v.png", name))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	graph.Render(chart.PNG, f)

	return name + ".png"
}
