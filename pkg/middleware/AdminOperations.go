package middleware

import (
	"CommunistsStatistic/pkg/DataBase"
	Fed "CommunistsStatistic/pkg/FedStatQueryHandler"
	imgMaker "CommunistsStatistic/pkg/charts_img_maker"
	"CommunistsStatistic/pkg/encharts_maker"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type DatasetsShort struct {
	Id   int
	Name string
}

type ChartsShort struct {
	Id   int
	Name string
}

type OutputChartAdmin struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Order       int    `json:"Order"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	HtmlChart   string `json:"HtmlChart"`
	AllDatasets []DatasetsShort
	Datasets    []encharts_maker.DataSets `json:"DataSets"`
}

type OutputViewAdmin struct {
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

type PassToImg struct {
	MainChartId int
	MainAxisId  int
	DataSets    []NestedDatasets
}

type NestedDatasets struct {
	Data map[int]int `json:"data"`
	Id   int         `json:"id"`
}

type SelectedDatasets []int

type SelectedCharts []int

func GetAllCharts() []string {
	return database.GetAllCharts()
}

func GetChartAdmin(id int) string {
	var ch encharts_maker.Chart
	var dts DatasetsShort
	fltrd := make([]DatasetsShort, 0)
	json.Unmarshal([]byte(database.GetChart(id)), &ch)
	html := encharts_maker.GenerateChart(ch)

	for _, v := range database.GetDatasets() {
		json.Unmarshal([]byte(v), &dts)
		fltrd = append(fltrd, dts)
	}

	out := OutputChartAdmin{Id: ch.Id, Order: ch.Order, Name: ch.Name, Title: ch.Title,
		Description: ch.Description, HtmlChart: string(html), AllDatasets: fltrd, Datasets: ch.DataSets}
	bytes, _ := json.Marshal(out)

	return string(bytes)
}

func CreateChart(name string, title string) error {
	return database.CreateChart(name, title)
}

func GetDataset(id int) []string {
	return database.GetDataset(id)
}

func GetDatasets() []string {
	return database.GetDatasets()
}

func SelectDatasets(data string, idStr string) error {
	var datasetsIds SelectedDatasets
	json.Unmarshal([]byte(data), &datasetsIds)
	id, _ := strconv.Atoi(idStr)

	return database.BoundDatasetToChart(datasetsIds, id)
}

func SetAxis(datasetId string, chartId string) error {
	dataset_id, _ := strconv.Atoi(datasetId)
	chart_id, _ := strconv.Atoi(chartId)
	return database.SetAxis(dataset_id, chart_id)
}

func UpdateChart(name string, order string, title string, description string, chartId string) error {
	chart_id, _ := strconv.Atoi(chartId)
	return database.UpdateChart(name, order, title, description, chart_id)
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

func GetViewAdmin(id string) string {
	view_id, _ := strconv.Atoi(id)
	separated := strings.Join(database.GetChartsShort(), ",")
	return fmt.Sprintf("{\"general_data\":%v, \"chars_shorts\": [%v]}", database.GetView(view_id), separated)
}

func UpdateView(name string, title string, seoDescription string, seoKeywords string, description string, id string) error {
	intId, _ := strconv.Atoi(id)
	return database.UpdateView(name, title, seoDescription, seoKeywords, description, intId)
}

func SetCharts(data string, idStr string) error {
	var chartsIds SelectedCharts
	json.Unmarshal([]byte(data), &chartsIds)
	id, _ := strconv.Atoi(idStr)

	return database.BoundChartsToView(chartsIds, id)
}

func SetMainChart(chartIdSrt string, viewIdSrt string) error {
	chartId, _ := strconv.Atoi(chartIdSrt)
	viewId, _ := strconv.Atoi(viewIdSrt)
	return database.SetMainChart(chartId, viewId)
}

func CreateView(name string, title string) error {
	return database.CreateView(name, title)
}

func CreateDataset(name string) error {
	return database.CreateDataset(name)
}

func DeleteDataset(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	return database.DeleteDataset(id)
}

func DeleteChart(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	return database.DeleteChart(id)
}

func DeleteView(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	return database.DeleteView(id)
}

func CreateImg(id string) error {
	var d PassToImg
	var values map[int]int

	view_id, _ := strconv.Atoi(id)
	data := database.GetView(view_id)
	json.Unmarshal([]byte(data), &d)

	data = database.GetChart(d.MainChartId)
	json.Unmarshal([]byte(data), &d)
	for _, v := range d.DataSets {
		if v.Id == d.MainAxisId {
			values = v.Data
		}
	}
	maker := filterDataForImage(values)
	name := maker.MakeChartImg()
	err := database.UpdateViewImgName(name, view_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func filterDataForImage(raw map[int]int) imgMaker.Chart {
	j := imgMaker.Chart{}
	for k, v := range raw {
		j.XValues = append(j.XValues, float64(k))
		j.YValues = append(j.YValues, float64(v))
	}
	return j
}
