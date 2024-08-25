package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type DataSet struct {
	ID      int
	Name    string
	Data    map[string]int
	RawData string
	RawUrl  string
}

type DataSets struct {
	Id   int
	Name string
}

func GetDataset(id int) []string {
	result := make([]string, 0)
	rows, err := conn.Query(context.Background(),
		fmt.Sprintf("SELECT id, name, data, coalesce(raw_data,''), coalesce(raw_url,'')  FROM dataset WHERE id = %d", id))
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id int
		var name, rawData, rawUrl string
		var data map[string]int
		err := rows.Scan(&id, &name, &data, &rawData, &rawUrl)
		arr := DataSet{id, name, data, rawData, rawUrl}
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

func GetDatasets() []string {
	result := make([]string, 0)
	rows, err := conn.Query(context.Background(),
		"SELECT id, coalesce(name, '') FROM dataset")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		arr := DataSets{id, name}
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

func UpdateDataset(mode string, name string, raw string, data string, id int) error {
	if mode == "raw_data" {
		_, err := conn.Query(context.Background(),
			fmt.Sprintf("UPDATE dataset SET (name, data, raw_data, raw_url) = ('%s', '%s', '%s', null) WHERE id = %d", name, data, raw, id))
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("%v", err))
		}

	} else {
		fmt.Println(data)
		_, err := conn.Query(context.Background(),
			fmt.Sprintf("UPDATE dataset SET (name, data, raw_data, raw_url) = ('%s', '%s', null, '%s') WHERE id = %d", name, data, raw, id))
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("%v", err))
		}

	}
	return nil
}

func CreateDataset(name string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO dataset(name) VALUES($1)", name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteDataset(id int) error {
	_, err := conn.Exec(context.Background(), "DELETE from dataset WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
