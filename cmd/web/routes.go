package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/privacy", app.loginInAdmin)
	mux.HandleFunc("/terms", app.adminPanel)
	mux.HandleFunc("/terms/charts", app.GetCharts)
	mux.HandleFunc("/terms/dataset/", app.GetDataset)
	mux.HandleFunc("/terms/datasets/update", app.UpdateDataset)
	mux.HandleFunc("/terms/chart/", app.GetChart)
	mux.HandleFunc("/terms/select_datasets", app.SelectDatasets)
	mux.HandleFunc("/terms/create_chart", app.CreateChart)
	mux.HandleFunc("/terms/set_axis", app.SetAxis)
	mux.HandleFunc("/terms/update_chart", app.UpdateChart)
	mux.HandleFunc("/terms/get_views", app.GetViews)
	mux.HandleFunc("/terms/get_view", app.GetView)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
