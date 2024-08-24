package middleware

import (
	db "CommunistsStatistic/DataBase"
	Fed "CommunistsStatistic/FedStatQueryHandler"
	"CommunistsStatistic/encharts_maker"
	"encoding/json"
	"strconv"
)

type DatasetsShort struct {
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

type SelectedDatasets []int

func GetCharts() []string {
	return db.GetCharts()
}

func GetChart(id int) string {
	var ch encharts_maker.Chart
	var dts DatasetsShort
	fltrd := make([]DatasetsShort, 0)
	json.Unmarshal([]byte(db.GetChart(id)), &ch)
	html := encharts_maker.GenerateChart(ch)

	for _, v := range db.GetDatasets() {
		json.Unmarshal([]byte(v), &dts)
		fltrd = append(fltrd, dts)
	}

	out := OutputData{Id: ch.Id, Name: ch.Name, Title: ch.Title, HtmlChart: string(html), AllDatasets: fltrd, Datasets: ch.DataSets}
	bytes, _ := json.Marshal(out)

	return string(bytes)
}

func CreateChart(name string, title string) error {
	return db.CreateChart(name, title)
}

func GetDataset(id int) []string {
	return db.GetDataset(id)
}

func SelectDatasets(data string, idstr string) error {
	var datasetsIds SelectedDatasets
	json.Unmarshal([]byte(data), &datasetsIds)
	id, _ := strconv.Atoi(idstr)

	return db.BoundDatasetToChart(datasetsIds, id)
}

func SetAxis(datasetId string, chartId string) error {
	dataset_id, _ := strconv.Atoi(datasetId)
	chart_id, _ := strconv.Atoi(chartId)
	return db.SetAxis(dataset_id, chart_id)
}

func UpdateChart(name string, title string, chartId string) error {
	chart_id, _ := strconv.Atoi(chartId)
	return db.UpdateChart(name, title, chart_id)
}

func UpdateDataset(mode string, name string, raw string, id int) error {
	if mode == "raw_data" {
		var data = raw
		return db.UpdateDataset(mode, name, raw, data, id)
	} else {
		data := Fed.GetStatData(raw)
		return db.UpdateDataset(mode, name, raw, data, id)
	}
}
