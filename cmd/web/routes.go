package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/privacy", app.loginInAdmin)
	mux.HandleFunc("/terms", app.adminPanel)
	mux.HandleFunc("/terms/charts", app.GetCharts)
	mux.HandleFunc("/terms/dataset/", app.GetDataset)
	mux.HandleFunc("/terms/datasets/create", app.CreateDataset)
	mux.HandleFunc("/terms/datasets/get_all", app.GetDatasets)
	mux.HandleFunc("/terms/datasets/update", app.UpdateDataset)
	mux.HandleFunc("/terms/datasets/delete", app.deleteDataset)
	mux.HandleFunc("/terms/chart/", app.GetChart)
	mux.HandleFunc("/terms/select_datasets", app.SelectDatasets)
	mux.HandleFunc("/terms/create_chart", app.CreateChart)
	mux.HandleFunc("/terms/set_axis", app.SetAxis)
	mux.HandleFunc("/terms/update_chart", app.UpdateChart)
	mux.HandleFunc("/terms/charts/delete", app.deleteChart)
	mux.HandleFunc("/terms/get_views", app.GetViews)
	mux.HandleFunc("/terms/get_view", app.GetView)
	mux.HandleFunc("/terms/update_view", app.UpdateView)
	mux.HandleFunc("/terms/views/set_charts", app.SetCharts)
	mux.HandleFunc("/terms/views/set_main_chart", app.SetMainChart)
	mux.HandleFunc("/terms/views/create_view", app.createView)
	mux.HandleFunc("/terms/views/delete", app.deleteView)
	mux.HandleFunc("/terms/views/create_img", app.createImg)
	mux.HandleFunc("/views/get_all", app.getViews)
	mux.HandleFunc("/views/get_view", app.getView)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
