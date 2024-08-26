package middleware

import (
	"CommunistsStatistic/cmd/DataBase"
	"CommunistsStatistic/cmd/encharts_maker"
	"encoding/json"
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
	//charts := make([]OutputChart, 0)
	var out_view OutputView
	var charts_ids []int

	id, _ := strconv.Atoi(idStr)
	raw_view := database.GetView(id)
	json.Unmarshal([]byte(raw_view), &out_view)
	for _, v := range out_view.BoundedCharts {
		charts_ids = append(charts_ids, v.Id)
	}

	out_view.BoundedCharts = []OutputChart{}
	for _, v := range charts_ids {
		out_view.BoundedCharts = append(out_view.BoundedCharts, getChart(v))
	}

	jdata, _ := json.Marshal(out_view)

	return jdata
}

func getChart(id int) OutputChart {
	var ch encharts_maker.Chart
	json.Unmarshal([]byte(database.GetChart(id)), &ch)
	html := encharts_maker.GenerateChart(ch)
	out := OutputChart{Id: ch.Id, Name: ch.Name, Title: ch.Title, HtmlChart: string(html)}

	return out
}
