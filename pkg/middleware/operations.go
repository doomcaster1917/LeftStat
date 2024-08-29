package middleware

import (
	"CommunistsStatistic/pkg/DataBase"
	"CommunistsStatistic/pkg/encharts_maker"
	"encoding/json"
	"html/template"
	"strconv"
)

type Views struct {
	Id      int
	Name    string
	Title   string
	ImgAddr string
}

type OutputChart struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Title       string `json:"Title"`
	HtmlChart   string `json:"HtmlChart"`
	AllDatasets []DatasetsShort
	Datasets    []encharts_maker.DataSets `json:"DataSets"`
}

type OutputView struct {
	Id             int    `json:"Id"`
	Name           string `json:"Name"`
	Title          string `json:"Title"`
	ImgAddr        string
	SeoDescription string
	SeoKeywords    string
	Description    string
	BoundedCharts  []OutputChart
}

func GetAllViews() []string {
	return database.GetViews()
}

func GetView(idStr string) []byte {

	var out_view OutputView
	var charts_ids []int
	var html OutputChart

	id, _ := strconv.Atoi(idStr)
	raw_view := database.GetView(id)
	json.Unmarshal([]byte(raw_view), &out_view)
	for _, v := range out_view.BoundedCharts {
		charts_ids = append(charts_ids, v.Id)
	}

	out_view.BoundedCharts = []OutputChart{}
	for _, v := range charts_ids {
		html = getChart(v)
		if len(html.HtmlChart) != 0 {
			out_view.BoundedCharts = append(out_view.BoundedCharts, html)
		} else {
			continue
		}
	}

	jdata, _ := json.Marshal(out_view)

	return jdata
}

func getChart(id int) OutputChart {
	var ch encharts_maker.Chart
	var html template.HTML
	json.Unmarshal([]byte(database.GetChart(id)), &ch)
	if len(ch.DataSets) > 0 && len(ch.DataSets[0].Data) != 0 {
		html = encharts_maker.GenerateChart(ch)
	} else {
		html = ""
	}
	out := OutputChart{Id: ch.Id, Name: ch.Name, Title: ch.Title, HtmlChart: string(html)}

	return out
}
