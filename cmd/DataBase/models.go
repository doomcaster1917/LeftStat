package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type tools interface {
	CreateConnection() *pgxpool.Pool
}

func RefreshModels() {
	conn := tools.CreateConnection(Db{})
	_, err := conn.Exec(context.Background(), "DROP TABLE IF EXISTS chart_view; DROP TABLE IF EXISTS dataset_chart; "+
		"DROP TABLE IF EXISTS chart; DROP TABLE IF EXISTS dataset; DROP TABLE IF EXISTS view;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed 1: %v\n", err)
	}

	_, err = conn.Exec(context.Background(), "CREATE TABLE view (id SERIAL PRIMARY KEY, name VARCHAR(250), title VARCHAR(250), img_addr VARCHAR(200), seo_description TEXT, seo_keywords TEXT, description TEXT)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed 4: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(context.Background(), "CREATE TABLE chart (id SERIAL PRIMARY KEY, "+
		"name VARCHAR(250), title VARCHAR(250))")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed 5: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), "CREATE TABLE dataset (id SERIAL PRIMARY KEY, name VARCHAR(250), data JSON, raw_data TEXT)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed 6: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), ""+
		"CREATE TABLE dataset_chart (dataset_id INTEGER REFERENCES dataset (id)"+
		"ON UPDATE CASCADE ON DELETE CASCADE,"+
		" chart_id INTEGER REFERENCES chart (id)"+
		"ON UPDATE CASCADE ON DELETE CASCADE, CONSTRAINT dataset_chart_pkey PRIMARY KEY (dataset_id, chart_id))")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), ""+
		"CREATE TABLE chart_view (chart_id INTEGER REFERENCES chart (id) "+
		"ON UPDATE CASCADE, view_id INTEGER REFERENCES view (id) "+
		"ON UPDATE CASCADE, CONSTRAINT chart_view_pkey PRIMARY KEY (chart_id, view_id))")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
