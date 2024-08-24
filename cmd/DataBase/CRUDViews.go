package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type View struct {
	ID               int
	Name             string
	Title            string
	ImgAddr          string
	Seo_descriptions string
	Seo_keywords     string
	Description      string
	BoundedCharts    []map[string]interface{}
}

type Views struct {
	ID      int
	Name    string
	Title   string
	ImgAddr string
}

func GetViews() []string {
	var result []string
	rows, err := conn.Query(context.Background(),
		"SELECT id, coalesce(name,''), coalesce(title,''), coalesce(img_addr, '') FROM view GROUP BY id")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name, title, imgAddr string
		err := rows.Scan(&id, &name, &title, &imgAddr)
		arr := Views{id, name, title, imgAddr}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(arr)
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
		"SELECT coalesce(v.id, 0), coalesce(v.name,''), coalesce(v.title,''), coalesce(v.img_addr, ''), coalesce(v.seo_description,''), coalesce(v.seo_keywords,''), coalesce(v.description,''), JSON_AGG(json_build_object('id', c.id, 'name', c.name, 'title', c.title)) as bounded_charts FROM chart_view LEFT JOIN view v ON v.id = chart_view.view_id LEFT JOIN chart c ON c.id = chart_view.chart_id WHERE v.id = $1 GROUP BY v.id", id)

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name, title, seoDescription, seoKeywords, description, imgAddr string
		var boundedCharts []map[string]interface{}
		err := rows.Scan(&id, &name, &title, &imgAddr, &seoDescription, &seoKeywords, &description, &boundedCharts)
		arr := View{id, name, title, imgAddr, seoDescription, seoKeywords, seoKeywords, boundedCharts}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(arr)
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result += string(bytes)
	}
	return result
}
