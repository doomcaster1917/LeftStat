package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Charts struct {
	Id       int
	Name     string
	Title    string
	DataSets []map[string]interface{}
}

type Chart struct {
	Id          int
	Order       int
	Name        string
	Title       string
	Description string
	MainAxisId  int
	DataSets    []map[string]interface{}
}

type ChartsShort struct {
	Id   int
	Name string
}

type ChartBounds struct {
	DatasetId   int
	DatasetName string
}

func GetAllCharts() []string {
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

func GetChartsShort() []string {
	result := make([]string, 0)
	rows, err := conn.Query(context.Background(),
		"SELECT coalesce(id, 0), coalesce(name,'') FROM chart GROUP BY id")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		arr := ChartsShort{id, name}
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
		"SELECT c.id, coalesce(c.return_order, 0), c.name, coalesce(c.title, ''), coalesce(c.description, ''), coalesce(c.main_axis_id, 0),"+
			" JSON_AGG(json_build_object('id', d.id, 'name', d.name, 'data', d.data)) "+
			"as dataset_name FROM dataset_chart "+
			"FULL JOIN chart c ON c.id = dataset_chart.chart_id "+
			"FULL JOIN dataset d ON d.id = dataset_chart.dataset_id "+
			"WHERE c.id = %d "+
			"GROUP BY c.id", id))

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id, order, mainAxisId int
		var name, title, description string
		var datasets []map[string]interface{}
		err := rows.Scan(&id, &order, &name, &title, &description, &mainAxisId, &datasets)
		arr := Chart{id, order, name, title, description, mainAxisId, datasets}
		if err != nil {
			fmt.Println(111, err)
		}
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result += string(bytes)
	}
	return result
}

func GetCharts(ids []string) string {
	var result []string
	params := "{" + strings.Join(ids, ",") + "}"
	rows, err := conn.Query(context.Background(),
		"SELECT c.id, coalesce(c.return_order, 0), c.name, coalesce(c.title, ''), coalesce(c.description, ''), coalesce(c.main_axis_id, 0),"+
			" JSON_AGG(json_build_object('id', d.id, 'name', d.name, 'data', d.data)) "+
			"as dataset_name FROM dataset_chart "+
			"FULL JOIN chart c ON c.id = dataset_chart.chart_id "+
			"FULL JOIN dataset d ON d.id = dataset_chart.dataset_id "+
			"WHERE c.id = ANY($1::int[]) "+
			"GROUP BY c.id ORDER BY c.return_order", params)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id, order, mainAxisId int
		var name, title, description string
		var datasets []map[string]interface{}
		err := rows.Scan(&id, &order, &name, &title, &description, &mainAxisId, &datasets)
		arr := Chart{id, order, name, title, description, mainAxisId, datasets}
		if err != nil {
			fmt.Println(111, err)
		}
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, string(bytes))
	}
	return fmt.Sprintf("[%s]", strings.Join(result, `,`))
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

func UpdateChart(name string, order string, title string, description string, id int) error {
	_, err := conn.Exec(context.Background(), "UPDATE chart SET (name, return_order, title, description) "+
		"= ($1, $2, $3, $4) WHERE id = $5", name, order, title, description, id)
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

func DeleteChart(id int) error {
	_, err := conn.Exec(context.Background(), "DELETE from chart WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
