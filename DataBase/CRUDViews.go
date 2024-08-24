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
	Seo_descriptions string
	Seo_keywords     string
	Description      string
	PicUrl           string
}

func addDataset() {

}

func GetViews() []byte {
	var result []byte
	rows, err := conn.Query(context.Background(),
		"SELECT p.full_name AS person_full_name, b.name AS bank_name"+
			"FROM person p"+
			"LEFT JOIN person_bank pb ON pb.person_id = p.id"+
			"LEFT JOIN bank b ON b.id = pb.bank_id;")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name, title, seoDescription, seoKeywords, description, picUrl string
		err := rows.Scan(&id, &name, &title, &seoDescription, &seoKeywords, &description, &picUrl)
		arr := View{id, name, title, seoDescription, seoKeywords, seoKeywords, picUrl}
		if err != nil {
			fmt.Println(err)
		}
		bytes, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, bytes...)

	}
	return result
}
