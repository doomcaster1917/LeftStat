package middleware

import (
	"CommunistsStatistic/cmd/DataBase"
	Fed "CommunistsStatistic/cmd/FedStatQueryHandler"
	"CommunistsStatistic/cmd/encharts_maker"
	"encoding/json"
	"fmt"
	"strconv"
)

type DatasetsShort struct {
	Id   int
	Name string
}

type ChartsShort struct {
	Id   int
	Name string
}

type OutputData struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Title       string `json:"Title"`
	HtmlChart   string `json:"HtmlChart"`
	AllDatasets []DatasetsShort
	Datasets    []encharts_maker.DataSets `json:"DataSets"`
}

type OutputView struct {
	Id             int
	Name           string
	Title          string
	ImgAddr        string
	SeoDescription string
	SeoKeywords    string
	Description    string
	BoundedCharts  string
	AllCharsShorts []ChartsShort
}

type SelectedDatasets []int

func GetCharts() []string {
	return database.GetCharts()
}

func GetChart(id int) string {
	var ch encharts_maker.Chart
	var dts DatasetsShort
	fltrd := make([]DatasetsShort, 0)
	json.Unmarshal([]byte(database.GetChart(id)), &ch)
	html := encharts_maker.GenerateChart(ch)

	for _, v := range database.GetDatasets() {
		json.Unmarshal([]byte(v), &dts)
		fltrd = append(fltrd, dts)
	}

	out := OutputData{Id: ch.Id, Name: ch.Name, Title: ch.Title, HtmlChart: string(html), AllDatasets: fltrd, Datasets: ch.DataSets}
	bytes, _ := json.Marshal(out)

	return string(bytes)
}

func CreateChart(name string, title string) error {
	return database.CreateChart(name, title)
}

func GetDataset(id int) []string {
	return database.GetDataset(id)
}

func SelectDatasets(data string, idstr string) error {
	var datasetsIds SelectedDatasets
	json.Unmarshal([]byte(data), &datasetsIds)
	id, _ := strconv.Atoi(idstr)

	return database.BoundDatasetToChart(datasetsIds, id)
}

func SetAxis(datasetId string, chartId string) error {
	dataset_id, _ := strconv.Atoi(datasetId)
	chart_id, _ := strconv.Atoi(chartId)
	return database.SetAxis(dataset_id, chart_id)
}

func UpdateChart(name string, title string, chartId string) error {
	chart_id, _ := strconv.Atoi(chartId)
	return database.UpdateChart(name, title, chart_id)
}

func UpdateDataset(mode string, name string, raw string, id int) error {
	if mode == "raw_data" {
		var data = raw
		return database.UpdateDataset(mode, name, raw, data, id)
	} else {
		data := Fed.GetStatData(raw)
		return database.UpdateDataset(mode, name, raw, data, id)
	}
}

func GetViews() []string {
	return database.GetViews()
}

func GetView(id string) string {
	//fltrd := make([]ChartsShort, 0)
	//
	view_id, _ := strconv.Atoi(id)
	//genData := database.GetView(view_id)
	//json.Unmarshal([]byte(genData), &fltrd)
	//out := OutputView{}
	return fmt.Sprintf("{\"general_data\":%v, \"chars_shors\": %v}", database.GetView(view_id), database.GetChartsShort())
}
