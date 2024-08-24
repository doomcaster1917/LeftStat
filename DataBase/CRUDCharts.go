package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Charts struct {
	Id       int
	Name     string
	Title    string
	DataSets []map[string]interface{}
}

type Chart struct {
	Id         int
	Name       string
	Title      string
	MainAxisId int
	DataSets   []map[string]interface{}
}

type ChartBounds struct {
	DatasetId   int
	DatasetName string
}

func GetCharts() []string {
	result := make([]string, 0)
	rows, err := conn.Query(context.Background(),
		"SELECT coalesce(c.id, 0), coalesce(c.name,''), coalesce(c.title,''), JSON_AGG(json_build_object('id', d.id, 'name', d.name)) "+
			"as dataset_name FROM dataset_chart"+
			" FULL OUTER JOIN chart c ON c.id = dataset_chart.chart_id"+
			" FULL OUTER JOIN dataset d ON d.id = dataset_chart.dataset_id GROUP BY c.id")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name, title string
		var datasets []map[string]interface{}
		err := rows.Scan(&id, &name, &title, &datasets)
		arr := Charts{id, name, title, datasets}
		if err != nil {
			fmt.Println(err)
		}
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, string(bytes))
	}
	return result
}

func GetChart(id int) string {
	var result string
	rows, err := conn.Query(context.Background(), fmt.Sprintf(
		"SELECT c.id, c.name, c.title, coalesce(c.main_axis_id, 0), JSON_AGG(json_build_object('id', d.id, 'name', d.name, 'data', d.data)) "+
			"as dataset_name FROM dataset_chart "+
			"FULL JOIN chart c ON c.id = dataset_chart.chart_id "+
			"FULL JOIN dataset d ON d.id = dataset_chart.dataset_id "+
			"WHERE c.id = %d"+
			" GROUP BY c.id", id))

	if err != nil {
		fmt.Println(222, err)
	}
	for rows.Next() {
		var id, mainAxisId int
		var name, title string
		var datasets []map[string]interface{}
		err := rows.Scan(&id, &name, &title, &mainAxisId, &datasets)
		arr := Chart{id, name, title, mainAxisId, datasets}
		if err != nil {
			fmt.Println(111, err)
		}
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(333, err)
			os.Exit(1)
		}
		result += string(bytes)
	}
	return result
}

func BoundDatasetToChart(datasets []int, id int) error {

	_, err := conn.Query(context.Background(), fmt.Sprintf("DELETE FROM dataset_chart WHERE chart_id = %d", id))
	if err != nil {
		fmt.Println(err)
		return errors.New(fmt.Sprintf("%v", err))
	}

	for _, dataset_id := range datasets {
		_, err = conn.Exec(context.Background(), "INSERT INTO dataset_chart(chart_id, dataset_id) VALUES($1, $2)", id, dataset_id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func CreateChart(name string, title string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO chart(name, title) VALUES($1, $2)", name, title)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateChart(name string, title string, id int) error {
	_, err := conn.Exec(context.Background(), "UPDATE chart SET (name, title) = ($1, $2) WHERE id = $3", name, title, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetAxis(datasetId int, chartId int) error {
	_, err := conn.Exec(context.Background(), "UPDATE chart SET main_axis_id = $1 WHERE id = $2", datasetId, chartId)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}
	return nil
}
