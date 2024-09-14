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
	Id             int
	Order          int
	Name           string
	Title          string
	Description    string
	SeoDescription string
	HtmlChart      string
	AllDatasets    []DatasetsShort
	Datasets       []encharts_maker.DataSets
}

type OutputView struct {
	Id             int
	Name           string
	Title          string
	ImgAddr        string
	SeoDescription string
	SeoKeywords    string
	Description    string
	BoundedCharts  []OutputChart
}

type DataSets struct {
	Data map[int]int
	Id   int
	Name string
}

type Chart struct {
	Id             int
	Order          int
	Name           string
	Title          string
	Description    string
	SeoDescription string
	DataSets       []DataSets
	MainAxisId     int
}

func GetAllViews() []string {
	return database.GetViews()
}

func GetView(idStr string) []byte {

	var out_view OutputView
	var charts_ids []string
	var raw_charts []encharts_maker.Chart
	//var chart encharts_maker.Chart

	id, _ := strconv.Atoi(idStr)
	raw_view := database.GetView(id)

	json.Unmarshal([]byte(raw_view), &out_view)
	for _, v := range out_view.BoundedCharts {
		charts_ids = append(charts_ids, strconv.Itoa(v.Id))
	}

	json.Unmarshal([]byte(database.GetCharts(charts_ids)), &raw_charts)

	out_view.BoundedCharts = []OutputChart{}
	for _, v := range raw_charts {
		html := getHtml(v)
		if len(html) != 0 {
			out_view.BoundedCharts = append(out_view.BoundedCharts,
				OutputChart{Id: v.Id, Order: v.Order, Name: v.Name, Title: v.Title,
					Description: v.Description, SeoDescription: v.SeoDescription, HtmlChart: html})
		} else {
			continue
		}
	}

	jdata, _ := json.Marshal(out_view)

	return jdata
}

func getHtml(ch encharts_maker.Chart) string {
	var html template.HTML
	if len(ch.DataSets) > 0 && len(ch.DataSets[0].Data) != 0 {
		html = encharts_maker.GenerateChart(ch)
	} else {
		html = ""
	}
	return string(html)
}
