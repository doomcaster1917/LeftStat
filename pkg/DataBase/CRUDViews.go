package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type View struct {
	Id             int
	Name           string
	Title          string
	Slug           string
	ImgAddr        string
	SeoDescription string
	SeoKeywords    string
	Description    string
	MainChartId    int
	BoundedCharts  []map[string]interface{}
}

type Views struct {
	Id            int
	Name          string
	Title         string
	Slug          string
	ImgAddr       string
	BoundedCharts []map[string]interface{}
}

func GetViews() []string {
	var result []string
	rows, err := conn.Query(context.Background(),
		"SELECT coalesce(v.id, 0), coalesce(v.name,''), coalesce(v.title,''), v.slug, coalesce(v.img_addr, ''), "+
			"JSON_AGG(json_build_object('id', c.id, 'name', c.name, 'title', c.title)) "+
			"as bounded_charts FROM chart_view "+
			"RIGHT JOIN view v ON v.id = chart_view.view_id "+
			"LEFT JOIN chart c ON c.id = chart_view.chart_id "+
			"GROUP BY v.id")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name, title, slug, imgAddr string
		var boundedCharts []map[string]interface{}
		err := rows.Scan(&id, &name, &title, &slug, &imgAddr, &boundedCharts)
		arr := Views{id, name, title, slug, imgAddr, boundedCharts}

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

func GetView(id int) string {
	var result string
	rows, err := conn.Query(context.Background(),
		"SELECT coalesce(v.id, 0), coalesce(v.name,''), coalesce(v.title,''), v.slug, coalesce(v.img_addr, ''), coalesce(v.seo_description,''), coalesce(v.seo_keywords,''), coalesce(v.description,''), coalesce(v.main_chart_id,0), "+
			"JSON_AGG(json_build_object('id', c.id, 'name', c.name, 'title', c.title)) as bounded_charts "+
			"FROM chart_view "+
			"FULL JOIN view v ON v.id = chart_view.view_id "+
			"FULL JOIN chart c ON c.id = chart_view.chart_id WHERE v.id = $1 GROUP BY v.id", id)

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id, mainChartId int
		var name, title, slug, seoDescription, seoKeywords, description, imgAddr string
		var boundedCharts []map[string]interface{}
		err := rows.Scan(&id, &name, &title, &slug, &imgAddr, &seoDescription, &seoKeywords, &description, &mainChartId, &boundedCharts)
		arr := View{id, name, title, slug, imgAddr, seoDescription, seoKeywords, description, mainChartId, boundedCharts}
		if err != nil {
			fmt.Println(err)
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

func UpdateView(name, title, slug, seoDescription, seoKeywords, description string, id int) error {
	_, err := conn.Exec(context.Background(),
		"UPDATE view SET (name, title, slug, seo_description, seo_keywords, description) = ($1, $2, $3, $4, $5, $6) WHERE id = $7", name, title, slug, seoDescription, seoKeywords, description, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateViewImgName(imgName string, id int) error {
	fmt.Println(imgName, id)
	_, err := conn.Exec(context.Background(),
		"UPDATE view SET img_addr = $1 WHERE id = $2", imgName, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func BoundChartsToView(chartsIds []int, viewId int) error {

	_, err := conn.Query(context.Background(), fmt.Sprintf("DELETE FROM chart_view WHERE view_id = %d", viewId))
	if err != nil {
		fmt.Println(err)
		return errors.New(fmt.Sprintf("%v", err))
	}

	for _, dataset_id := range chartsIds {
		_, err = conn.Exec(context.Background(), "INSERT INTO chart_view(view_id, chart_id) VALUES($1, $2)", viewId, dataset_id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func SetMainChart(chartId int, viewId int) error {
	_, err := conn.Exec(context.Background(), "UPDATE view SET main_chart_id = $1 WHERE id = $2", chartId, viewId)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	return nil
}

func CreateView(name string, title string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO view(name, title) VALUES($1, $2)", name, title)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteView(id int) error {
	_, err := conn.Exec(context.Background(), "DELETE from view WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
