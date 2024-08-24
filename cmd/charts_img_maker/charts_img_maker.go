package charts_img_maker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wcharczuk/go-chart/v2"
	"os"
)

type ChartImager interface {
	makeChartImg() *os.File
}

type Chart struct {
	XValues []float64
	YValues []float64
}

func (c Chart) main() *os.File {

	graph := chart.Chart{
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
	f, _ := os.Create(fmt.Sprintf("static/%v.png", id.String()))
	defer f.Close()
	graph.Render(chart.PNG, f)

	return f
}
